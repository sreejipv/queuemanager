package main

import (
	"errors"
	"sync"
)

type Queue struct {
	items []string
	mutex sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		items: []string{},
	}
}

func (q *Queue) Enqueue(item string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (string, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.items) == 0 {
		return "", errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) List() []string {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.items
}
