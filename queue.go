package main

import "sync"

type ItemQ interface{}

type Queue struct {
	items []ItemQ
	mutex sync.Mutex
}

func (queue *Queue) Enqueue(item ItemQ) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.items = append(queue.items, item)
}

func (queue *Queue) Dequeue() ItemQ {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	if len(queue.items) == 0 {
		return nil
	}
	item := queue.items[0]
	queue.items = queue.items[1:]
	return item
}
