package main

import (
	"advent/pkg/line"
	"fmt"
	"log"
	"os"
)

const (
	size = 3
)

type Window struct {
	size int
	ns   []int
}

func NewWindow(n int) *Window {
	return &Window{
		size: n,
		ns:   make([]int, 0, n),
	}
}

func (w *Window) RunningSum(i int) (int, bool) {
	if n := len(w.ns); n >= w.size {
		w.ns = w.ns[1:]
	}
	w.ns = append(w.ns, i)
	if n := len(w.ns); n < w.size {
		return 0, false
	}
	sum := 0
	for _, i := range w.ns {
		sum += i
	}
	return sum, true
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <input>", os.Args[0])
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lastsum := -1
	n := 0
	w := NewWindow(size)
	for i := range line.StreamInt(f) {
		sum, ok := w.RunningSum(i)
		if !ok {
			continue
		}
		if lastsum < sum && lastsum >= 0 {
			n++
		}
		lastsum = sum
	}
	fmt.Println(n)
}
