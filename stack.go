package main

import "sync"

// Item s
type Item interface{}

//Stack s
type Stack struct {
	items []Item
	mutex sync.Mutex
}

//Push -> insert new item at the top of Stack
func (stack *Stack) Push(item Item) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = append(stack.items, item)
}

// Pop -> Remove top item from stack
func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	sizeItem := len(stack.items)
	if sizeItem == 0 {
		return nil
	}
	lastItem := stack.items[sizeItem-1]
	stack.items = stack.items[:sizeItem-1]
	return lastItem
}

//IsEmpty -> check whether stack empty or not
func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items) == 0
}

//Reset -> remove all items from stack
func (stack *Stack) Reset() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = nil
}

//Dump -> get all item from stack
func (stack *Stack) Dump() []Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	newItems := make([]Item, len(stack.items))
	copy(newItems, stack.items)
	return newItems
}

// Peek -> Returns the top item in the stack.
func (stack *Stack) Peek() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if len(stack.items) == 0 {
		return nil
	}
	return stack.items[len(stack.items)-1]
}
