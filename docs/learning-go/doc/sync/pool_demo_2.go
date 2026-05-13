package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID  uint
	Job func()
}

type Pool struct {
	taskQueue chan Task
	wg        sync.WaitGroup
}

func NewPool(numWorkers int) *Pool {
	p := &Pool{taskQueue: make(chan Task)}

	p.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go p.worker()
	}

	return p
}

func (p *Pool) AddTask(task Task) {
	p.taskQueue <- task
}

func (p *Pool) Wait() {
	close(p.taskQueue)
	p.wg.Wait()
}

func (p *Pool) worker() {
	for task := range p.taskQueue {
		fmt.Printf("worker %d started task %d\n", task.ID, task.ID)
		task.Job()
		fmt.Printf("worker %d finished task %d\n", task.ID, task.ID)
	}
	p.wg.Done()
}

func PoolNote2() {
	// 初始化协程池
	pool := NewPool(4)

	// 添加任务到协程池
	for i := 0; i < 20; i++ {
		taskID := uint(i)
		task := Task{
			ID: taskID,
			Job: func() {
				fmt.Printf("Task %d is running\n", taskID)
			},
		}
		pool.AddTask(task)
	}

	// 等待所有任务完成
	pool.Wait()
}
