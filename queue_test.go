package go_structures

import (
	"testing"
)

func TestSingleEnqueue(t *testing.T) {
	queue := new(Queue)

	checkQueueSize(queue, t, 0)

	queue.Enqueue(42)

	checkQueueSize(queue, t, 1)

	checkFirstAndLastElementInt(queue, t, 42, 42)
}

func TestMultipleEnqueue(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue(27)
	queue.Enqueue(30)

	checkQueueSize(queue, t, 2)

	checkFirstAndLastElementInt(queue, t, 27, 30)
}

func TestMultiTypeEnqueue(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue(1)
	queue.Enqueue("another element")

	checkQueueSize(queue, t, 2)

	if queue.first.value.(int) != 1 {
		t.Errorf("First Element does not match: should be %d but is %d",
			1, queue.first.value.(int))
	}

	if queue.last.value.(string) != "another element" {
		t.Errorf("Last Element does not match: should be %s but is %s",
			"another element", queue.first.value.(string))
	}
}

func TestSingleDequeue(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue(1)

	value, err := queue.Dequeue()
	checkError("Queue", err, t)

	checkQueueSize(queue, t, 0)

	if value.(int) != 1 {
		t.Errorf("Element does not match: should be %d but is %d",
			1, value.(int))
	}

	if queue.first != nil {
		t.Errorf("First Element does not match: should be %T but is %T",
			nil, queue.first)
	}

	if queue.last != nil {
		t.Errorf("Last Element does not match: should be %T but is %T",
			nil, queue.first)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	queue := new(Queue)

	checkQueueSize(queue, t, 0)

	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Empty Queue did not return an error")
	}
}

func TestDequeueMultiple(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	checkQueueSize(queue, t, 3)

	checkFirstAndLastElementInt(queue, t, 1, 3)

	value, err := queue.Dequeue()
	checkError("Queue", err, t)

	checkQueueSize(queue, t, 2)

	checkValueInt(value.(int), 1, t)

	checkFirstAndLastElementInt(queue, t, 2, 3)

	value, err = queue.Dequeue()
	checkError("Queue", err, t)

	checkQueueSize(queue, t, 1)

	checkValueInt(value.(int), 2, t)

	checkFirstAndLastElementInt(queue, t, 3, 3)
}

func TestQueueSize(t *testing.T) {
	queue := new(Queue)

	checkQueueSize(queue, t, 0)

	queue.size = 1

	checkQueueSize(queue, t, 1)

	queue.size = 27

	checkQueueSize(queue, t, 27)
}

func checkQueueSize(q *Queue, t *testing.T, size int) {
	if q.Size() != size {
		t.Errorf("Queue size should be %d but is %d", size, q.Size())
	}
}

func checkFirstAndLastElementInt(q *Queue, t *testing.T, first int, last int) {
	if q.first.value.(int) != first {
		t.Errorf("First Element does not match: should be %d but is %d",
			first, q.first.value.(int))
	}

	if q.last.value.(int) != last {
		t.Errorf("Last Element does not match: should be %d but is %d",
			last, q.first.value.(int))
	}
}

func checkError(s string, err error, t *testing.T) {
	if err != nil {
		t.Errorf("%s returned an error: %s", s, err)
	}
}

func checkValueInt(is int, should int, t *testing.T) {
	if is != should {
		t.Errorf("Element does not match: should be %d but is %d",
			should, is)
	}
}
