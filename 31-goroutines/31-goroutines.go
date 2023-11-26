package main

import (
	"fmt"
	"time"
)

func sleepyGopher(i int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore", i, "...")
}

func main() {
	go sleepyGopher(-1)
	time.Sleep(4 * time.Second)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
