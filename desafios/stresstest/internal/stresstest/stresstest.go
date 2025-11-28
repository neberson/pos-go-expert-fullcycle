package stresstest

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCount map[int]int
	Elapsed     time.Duration
}

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

func Run(cfg Config) Result {
	start := time.Now()
	statusCount := make(map[int]int)
	var statusMu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, cfg.Concurrency)

	for i := 0; i < cfg.Requests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			resp, err := http.Get(cfg.URL)
			if err != nil {
				statusMu.Lock()
				statusCount[0]++
				statusMu.Unlock()
				return
			}
			statusMu.Lock()
			statusCount[resp.StatusCode]++
			statusMu.Unlock()
			resp.Body.Close()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	return Result{
		StatusCount: statusCount,
		Elapsed:     elapsed,
	}
}
