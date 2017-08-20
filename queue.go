// The MIT License (MIT)
//
// Copyright (c) 2017 Fabian Wenzelmann
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package goqueue

type BenchQueue interface {
	Push(val int)
	Pop() int
	Empty() bool
}

type SimpleSliceQueue struct {
	elements []int // The elements, we don't return a new slice but overwrite this slice
}

func NewSimpleSliceQueue(initialCapacity int) *SimpleSliceQueue {
	return &SimpleSliceQueue{make([]int, 0, initialCapacity)}
}

func (q *SimpleSliceQueue) Push(val int) {
	q.elements = append(q.elements, val)
}

func (q *SimpleSliceQueue) Pop() int {
	val := q.elements[0]
	q.elements = q.elements[1:]
	return val
}

func (q *SimpleSliceQueue) Empty() bool {
	return len(q.elements) == 0
}

type node struct {
	succ  *node
	value int
}

type LinkedQueue struct {
	// first and last are the first and last node in the queue,
	// either they're both nil or they both are not nil
	first, last *node
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{nil, nil}
}

func (q *LinkedQueue) Push(val int) {
	// first check if queue is empty, if so create new node
	if q.first == nil {
		n := &node{succ: nil,
			value: val}
		q.first = n
		q.last = n
	} else {
		// we have a node, create a new one and set the successor of the last
		// node to this new node
		q.last.succ = &node{succ: nil,
			value: val}
		// update q.last
		q.last = q.last.succ
	}
}

func (q *LinkedQueue) Pop() int {
	// let it panic if there is no element
	// get the first value
	res := q.first.value
	// update first
	q.first = q.first.succ
	// check if the queue became empty, in this case set the last element
	// to nil as well
	if q.first == nil {
		q.last = nil
	}
	return res
}

func (q *LinkedQueue) Empty() bool {
	return q.first == nil
}

// Adjusted version of https://gist.github.com/moraes/2141121 by
// moraes on github: https://gist.github.com/moraes
type ExtendableQueue struct {
	nodes []int
	size  int
	head  int
	tail  int
	count int
}

func NewExtendableQueue(size int) *ExtendableQueue {
	return &ExtendableQueue{nodes: make([]int, size),
		size: size}
}

func (q *ExtendableQueue) Pop() int {
	res := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return res
}

func (q *ExtendableQueue) Push(val int) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]int, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = val
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *ExtendableQueue) Empty() bool {
	return q.count == 0
}
