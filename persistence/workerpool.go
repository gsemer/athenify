package persistence

import (
	"athenify/domain"
	"log"
	"sync"
)

type WorkerPool struct {
	workers int
	Jobs    chan domain.Job
	Results chan domain.Result
	wg      *sync.WaitGroup
}

func NewWorkerPool(workers int, jobs chan domain.Job, results chan domain.Result, wg *sync.WaitGroup) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		Jobs:    jobs,
		Results: results,
		wg:      wg,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for job := range wp.Jobs {
		result := job.Process()
		if result.Error != nil {
			log.Println(result.Error.Error())
		}
		wp.Results <- result
	}
}
