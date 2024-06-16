package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fibonacci(channel chan int, n int, wg *sync.WaitGroup) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
	wg.Done()
}

func receiver(channel chan int, wg *sync.WaitGroup) {
	for i := range channel {
		fmt.Println("Listened value : ", i)
	}
	wg.Done()
}

func bufferedChannelDemo() {
	channel := make(chan int, 10)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fibonacci(channel, 10, wg)
	go receiver(channel, wg)
	wg.Wait()
}

func doWork(t time.Duration, ch chan string) {
	fmt.Println("doing some work")
	time.Sleep(t)
	fmt.Println("Work done")
	ch <- fmt.Sprintf("Work done : %d", rand.Intn(100))
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("Started executing goroutine 1")
		time.Sleep(time.Second * 60)
		fmt.Println("Finished executing goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Started executing goroutine 2")
		time.Sleep(time.Second * 10)
		fmt.Println("Finished executing goroutine 2")
		wg.Done()
	}()

	wg.Wait()
}
