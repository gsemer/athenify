package persistence

import (
	"athenify/domain"
	"sync"
)

type WorkerPool struct {
	workers int
	Jobs    chan domain.Job
	wg      *sync.WaitGroup
}

func NewWorkerPool(workers int, jobs chan domain.Job, wg *sync.WaitGroup) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		Jobs:    jobs,
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
		job.Process()
	}
}
