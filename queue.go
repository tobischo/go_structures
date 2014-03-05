package go_structures

import (
	"errors"
)

// Queue which is exposed to be used by other packages
type Queue struct {
	first *element
	last  *element
	size  int
}

// Adds any value to the queue
func (queue *Queue) Enqueue(value interface{}) {
	e := &element{}
	e.value = value

	queue.last.next = e
	queue.last = e

	queue.size++
}

// Removes the first element from the queue and returns it
// Returns an error if the queue is empty
func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.size == 0 {
		return nil, errors.New("Queue empty")
	}

	value := queue.first.value
	queue.first = queue.first.next
	queue.size--
	return value, nil
}

// Returns the current size of the queue
func (queue *Queue) Size() int {
	return queue.size
}
