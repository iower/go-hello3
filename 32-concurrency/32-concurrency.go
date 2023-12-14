package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func worker() {
	n := 0
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
		}
	}
}

func main() {
	mu.Lock()
	defer mu.Unlock()

	go worker()
	time.Sleep(100)
}
