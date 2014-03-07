package go_structures

import (
	"testing"
)

func TestSinglePush(t *testing.T) {
	stack := new(Stack)

	checkStackSize(stack, t, 0)

	stack.Push(12)

	checkStackSize(stack, t, 1)

	checkTopValueInt(stack, t, 12)
}

func TestMultiplePush(t *testing.T) {
	stack := new(Stack)
	stack.Push(28)
	stack.Push(30)

	checkStackSize(stack, t, 2)

	checkTopValueInt(stack, t, 30)
}

func TestMultiTypePush(t *testing.T) {
	stack := new(Stack)
	stack.Push(28)
	value := "element"
	stack.Push(value)

	checkStackSize(stack, t, 2)

	if stack.top.value.(string) != value {
		t.Errorf("Top Element does not match: should be %s but is %s",
			value, stack.top.value.(string))
	}
}

func TestSinglePop(t *testing.T) {
	stack := new(Stack)
	stack.Push(12)

	value, err := stack.Pop()
	if err != nil {
		t.Errorf("Stack returned an error: %s", err)
	}

	if value.(int) != 12 {
		t.Errorf("Element does not match: should be %d but is %d",
			12, value.(int))
	}

	if stack.top != nil {
		t.Errorf("Top Element does not match: should be %T but is %T",
			nil, stack.top)
	}
}

func TestPopEmptyStack(t *testing.T) {
	stack := new(Stack)

	checkStackSize(stack, t, 0)

	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Empty Stack did not return an error")
	}
}

func TestPopMultiple(t *testing.T) {
	stack := new(Stack)
	stack.Push(12)
	stack.Push(13)
	stack.Push(27)

	checkStackSize(stack, t, 3)

	checkTopValueInt(stack, t, 27)

	value, err := stack.Pop()
	checkError("Stack", err, t)

	checkStackSize(stack, t, 2)

	checkValueInt(value.(int), 27, t)

	checkTopValueInt(stack, t, 13)

	value, err = stack.Pop()
	checkError("Stack", err, t)

	checkStackSize(stack, t, 1)

	checkValueInt(value.(int), 13, t)

	checkTopValueInt(stack, t, 12)
}

func TestStackSize(t *testing.T) {
	stack := new(Stack)

	checkStackSize(stack, t, 0)

	stack.size = 1

	checkStackSize(stack, t, 1)

	stack.size = 27

	checkStackSize(stack, t, 27)
}

func checkTopValueInt(s *Stack, t *testing.T, value int) {
	if s.top.value.(int) != value {
		t.Errorf("Top Element does not match: should be %d but is %d",
			value, s.top.value.(int))
	}
}

func checkStackSize(s *Stack, t *testing.T, size int) {
	if s.Size() != size {
		t.Errorf("Stack size should be %d but is %d", size, s.Size())
	}
}
