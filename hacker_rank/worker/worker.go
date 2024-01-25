package worker

func CreateWorkers[In, Out any](num int, jobs <-chan In, results chan<- Out, f func(arg In) Out) {
	for i := 0; i < num; i++ {
		go worker(i, jobs, results, f)
	}
}

func worker[In, Out any](id int, jobs <-chan In, results chan<- Out, f func(arg In) Out) {
	for job := range jobs {
		results <- f(job)
	}
}
