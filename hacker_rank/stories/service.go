package stories

import (
	"fmt"
	"gophercises/hacker_rank/client"
	"gophercises/hacker_rank/worker"
	"sync"
	"time"
)

type Stories struct {
	client     *client.HackerRankClient
	cache      []client.Item
	mu         sync.Mutex
	expiration time.Time
	duration   time.Duration
	count      int
}

func NewStories(hrClient *client.HackerRankClient, funOpts ...func(stories *Stories)) *Stories {
	s := Stories{client: hrClient}
	for _, f := range funOpts {
		f(&s)
	}
	return &s
}

func WithDefaultParams(stories *Stories) {
	stories.mu = sync.Mutex{}
	stories.duration = 2 * time.Second
	stories.count = 30
}

func (s *Stories) GetCachedStories() ([]client.Item, error) {
	s.mu.Lock()
	if time.Now().Sub(s.expiration) < 0 {
		s.mu.Unlock()
		return s.cache, nil
	}
	go func() {
		defer s.mu.Unlock()
		stories, err := s.fetchStories()
		if err != nil {
			fmt.Println("Error fetching stories", err)
		}
		s.cache = stories
		s.expiration = time.Now().Add(s.duration)
	}()
	return s.cache, nil //might be stale
}

func (s *Stories) fetchStories() ([]client.Item, error) {
	startTime := time.Now()
	ids, err := s.client.GetTopStories()
	if err != nil {
		return nil, err
	}
	res := make([]client.Item, 0)
	start := 0
	for len(res) < s.count && start < len(ids) {
		toFetch := (s.count - len(res)) * 5 / 4
		start += toFetch
		fetchedIds := ids[start:min(start+toFetch, len(ids))]
		res = append(s.fetchStoriesConcurrentlyById(fetchedIds))
	}
	fmt.Printf("Fetching executed in the background: %v \n", time.Now().Sub(startTime))
	return res[0:s.count], nil
}

func (s *Stories) fetchStoriesConcurrentlyById(ids []int) []client.Item {
	argsCh, outCh := make(chan int, len(ids)), make(chan client.Item, len(ids))
	worker.CreateWorkers(6, argsCh, outCh, s.client.GetItem)
	for _, id := range ids {
		argsCh <- id
	}
	close(argsCh)
	res := make([]client.Item, 0)
	for range ids {
		res = append(res, <-outCh)
	}
	return res
}
