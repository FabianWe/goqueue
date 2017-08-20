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

package tests

import (
	"testing"

	"github.com/FabianWe/goqueue"
)

// a global variable, we set it in the tests so the compiler does not optimize
// or benchmarks away...

var dummyQueue goqueue.BenchQueue
var dummyInt int

func BenchmarkSimpleSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := goqueue.NewSimpleSliceQueue(50000)
		runBench(q)
	}
}

func BenchmarkLinked(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := goqueue.NewLinkedQueue()
		runBench(q)
	}
}

func BenchmarkExtendable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := goqueue.NewExtendableQueue(50000)
		runBench(q)
	}
}

func BenchmarkAlternateSimpleSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := goqueue.NewSimpleSliceQueue(50000)
		runAlternateBench(q)
	}
}

func BenchmarkAlternateLinked(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := goqueue.NewLinkedQueue()
		runAlternateBench(q)
	}
}

func BenchmarkAlternateExtendable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := goqueue.NewExtendableQueue(50000)
		runAlternateBench(q)
	}
}

func runBench(q goqueue.BenchQueue) {
	// first add 100,000 elements
	// then remove 50,000 elements
	// then add 75,000 elements
	for i := 0; i < 100000; i++ {
		q.Push(i)
	}
	for i := 0; i < 50000; i++ {
		dummyInt = q.Pop()
	}
	for i := 0; i < 75000; i++ {
		q.Push(i)
	}
	dummyQueue = q
}

func runAlternateBench(q goqueue.BenchQueue) {
	q.Push(0)
	// do 250,000 times the following: get an element, add two new elements
	for i := 0; i < 250000; i++ {
		x := q.Pop()
		q.Push(x)
		q.Push(x + 1)
	}
	dummyQueue = q
}
