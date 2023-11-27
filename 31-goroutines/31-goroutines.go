package main

import (
	"fmt"
	"time"
)

func sleepyGopher(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("...snore", i, "...")
}

func sleepyGopherCh(i int, c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("...snore", i, "...")
	c <- i
}

func main() {
	go sleepyGopher(-1)
	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(2 * time.Second)

	fmt.Println("-----------")

	c := make(chan int)

	// deadlock
	// c <- 99
	// r := <-c
	// fmt.Println(r)

	for i := 0; i < 5; i++ {
		go sleepyGopherCh(i, c)
	}
	for i := 0; i < 5; i++ {
		routineId := <-c
		fmt.Println("id", routineId, "has finished")
	}

	fmt.Println("-----------")

	// select
	for i := 0; i < 5; i++ {
		c := make(chan int)
		go sleepyGopherCh(i, c)
		timeout := time.After(1000 * time.Millisecond)
		select {
		case routineId := <-c:
			fmt.Println("id", routineId, "has finished")
		case <-timeout:
			fmt.Println("timeout", i)
		}
	}
	time.Sleep(5 * time.Second)

	// empty select = wait
	// select {}
}
