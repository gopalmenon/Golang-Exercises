package queueusing2stacks

import "golangexercises/stacksandqueues"

type MyQueue struct {
	enqueueStack, dequeueStack *stacksandqueues.Stack
}

func (m *MyQueue) init() {
	m.enqueueStack = stacksandqueues.New()
	m.dequeueStack = stacksandqueues.New()
}

func New() *MyQueue {
	m := new(MyQueue)
	m.init()
	return m
}

func (m *MyQueue) Enqueue(v interface{}) {

	for v := m.dequeueStack.Pop(); v != nil; v = m.dequeueStack.Pop() {
		m.enqueueStack.Push(v)
	}
	m.enqueueStack.Push(v)

}

func (m *MyQueue) Dequeue() interface{} {

	for v := m.enqueueStack.Pop(); v != nil; v = m.enqueueStack.Pop() {
		m.dequeueStack.Push(v)
	}
	return m.dequeueStack.Pop()
}
