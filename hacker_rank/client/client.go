package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HackerRankClient struct {
	apiBase string
}

func WithDefaultApiBase() func(client *HackerRankClient) {
	return func(c *HackerRankClient) {
		c.apiBase = "https://hacker-news.firebaseio.com/v0"
	}
}

func NewHackerRankClient(opts ...func(client *HackerRankClient)) *HackerRankClient {
	var c HackerRankClient
	for _, f := range opts {
		f(&c)
	}
	return &c
}

func (c *HackerRankClient) GetTopStories() ([]int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/newstories.json", c.apiBase))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status was not succesfull: %s", resp.Status)
	}
	stories := make([]int, 0)
	err = json.NewDecoder(resp.Body).Decode(&stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (c *HackerRankClient) GetItem(id int) Item {
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apiBase, id))
	if err != nil {
		fmt.Println(err)
		return Item{}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("response status was not succesfull: %s", resp.Status)
		return Item{}
	}
	var i Item
	err = json.NewDecoder(resp.Body).Decode(&i)
	if err != nil {
		fmt.Printf("Cound not decode with id: %s", resp.Status)
		return Item{}
	}
	return i
}

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`

	// Only one of these should exist
	Text string `json:"text"`
	URL  string `json:"url"`
}
