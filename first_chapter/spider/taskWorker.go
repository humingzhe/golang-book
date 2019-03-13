package main

import (
	"fmt"
	"sync"
	"time"
)

type taskPool struct {
	queue   chan func()
	wait    *sync.WaitGroup
	workers int
}

func newTaskPool(workers int) *taskPool {
	return &taskPool{
		queue:   make(chan func()),
		wait:    new(sync.WaitGroup),
		workers: workers,
	}
}
// 执行Task
func (t *taskPool) work() {
	defer t.wait.Done()
	for task := range t.queue {
		task()
	}
}

func (t *taskPool) Start() {
	for i := 0; i < t.workers; i++ {
		t.wait.Add(1)
		go t.work()
	}
}

func (t *taskPool) Submit(task func()) {
	t.queue <- task
}

func (t *taskPool) Stop() {
	close(t.queue)
}

func (t *taskPool) Wait() {
	t.wait.Wait()
}

func main () {
	pool := newTaskPool(5)
	pool.Start()
	lll := []int{1,2,3,4,5,6,7}
	for _, l := range lll {
		res_l := l
		pool.Submit(func() {
			time.Sleep(time.Duration(l) * time.Second)
			fmt.Println(res_l)
		})
	}
	pool.Stop()
	pool.Wait()
}
