package domain

import (
	"log"
)

type Job interface {
	Process() Result
}

type Result struct {
	Body  string
	Error error
}

type WorkerPool struct {
	workers int
	jobs    chan Job
	result  chan Result
}

func NewWorkerPool(workers int, jobs chan Job) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		jobs:    jobs,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	for job := range wp.jobs {
		if result := job.Process(); result.Error != nil {
			log.Println(result.Error.Error())
		}
	}
}
