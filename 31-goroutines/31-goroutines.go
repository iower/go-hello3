package main

import (
	"fmt"
	"math/rand"
	"strings"
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

func sourceChan(downstream chan string) {
	for _, v := range []string{"str1", "str2", "str3"} {
		downstream <- v
	}
	downstream <- ""
}

func sourceChan2(downstream chan string) {
	for _, v := range []string{"str1", "str2", "str3"} {
		downstream <- v
	}
	close(downstream)
}

func filterChan(upstream, downstream chan string) {
	for {
		item := <-upstream
		if !strings.Contains(item, "2") {
			downstream <- item
		}
		if item == "" {
			return
		}
	}
}

func filterChan2(upstream, downstream chan string) {
	for {
		item, open := <-upstream
		if !open {
			close(downstream)
			return
		}
		if !strings.Contains(item, "2") {
			downstream <- item
		}
	}
}

func filterChan3(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "2") {
			downstream <- item
		}
	}
	close(downstream)
}

func printChan(upstream chan string) {
	for {
		item := <-upstream
		if item == "" {
			return
		}
		fmt.Println(item)
	}
}

func pringChan2(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
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

	// pipeline

	c0 := make(chan string)
	c1 := make(chan string)
	go sourceChan(c0)
	go filterChan(c0, c1)
	printChan(c1)

	// close channels

	chan1 := make(chan string)
	chan2 := make(chan string)
	go genChan(chan1)
	go memChan(chan1, chan2)
	printChan(chan2)
}

func genChan(downstream chan string) {
	for i := 0; i < 10; i++ {
		downstream <- "val-" + fmt.Sprintf("%v", rand.Intn(3))
	}
	close(downstream)
}

func memChan(upstream, downstream chan string) {
	var prev string
	for val := range upstream {
		if val != prev {
			prev = val
			downstream <- val
		}
	}
	close(downstream)
}

func words(upstream, downstream chan string) {
	for str := range upstream {
		for _, word := range strings.Fields(str) {
			downstream <- word
		}
	}
	close(downstream)
}
