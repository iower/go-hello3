package main

import (
	"fmt"
	"image"
	"log"
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
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current pos", pos)
			next = time.After(time.Second)
		}
	}
}

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 1
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
				log.Printf("new direction %v", direction)
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
				log.Printf("new direction %v", direction)
			case start:
				speed = 1
				log.Println("start")
			case stop:
				speed = 0
				log.Println("stop")
			}
		case <-nextMove:
			pos = pos.Add(direction.Mul(speed))
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func main() {
	mu.Lock()
	defer mu.Unlock()

	go worker()

	r := NewRoverDriver()
	time.Sleep(time.Second)
	r.Start()
	time.Sleep(time.Second)
	r.Left()
	time.Sleep(time.Second)
	r.Right()
	time.Sleep(time.Second)
	r.Stop()
	time.Sleep(time.Second)
	r.Right()
	time.Sleep(time.Second)
}
