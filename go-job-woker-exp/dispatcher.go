package main

import (
    "fmt"
    "log"
)

type Dispatcher struct {
    maxWorkers int
    WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
    pool := make(chan chan Job, maxWorkers)
    return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
    for i:= 0; i < d.maxWorkers; i++ {
        worker := NewWorker(d.WorkerPool)
        worker.Start()
    }

    go d.dispatch()
}

func (d * Dispatcher) dispatch() {
    fmt.Println("Worker que dispatcher started...")
    for {
        select {
        case job := <-JobQueue:
            log.Printf("a dispatcher request received")
            go func(job Job) {
                // try to obtain a worker job channel that is available.
                // this will block until a worker is idle
                jobChannel := <-d.WorkerPool

                // dispatch the job to the worker job channel
                jobChannel <- job
            }(job)

        }
    }
}
