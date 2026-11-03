package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		time.Sleep(100 * time.Millisecond)
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		// The select statement lets a goroutine wait on multiple communication operations
		// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	// c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	// defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime.
	// go ...
	// go say("world")
	// say("hello")

	// Channels are a typed conduit through which you can send and receive values with the channel operator, <-
	// ch <- v    // Send v to channel ch.
	// v := <-ch  // Receive from ch, and assign value to v.
	// (The data flows in the direction of the arrow.)
	s := []int{7, 2, 8, -9, 4, 0}

	// Like maps and slices, channels must be created before use
	// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Range and Close
	cb := make(chan int, 10)
	go fibonacci(cap(cb), cb)
	for i := range cb {
		fmt.Println(i)
	}

	fmt.Println("==== SELECT ====")
	cc := make(chan int)
	quit := make(chan int)
	go func() {
		for range 15 {
			fmt.Println(<-cc)
		}
		time.Sleep(2 * time.Second)
		fmt.Println(<-cc)
		quit <- 0
	}()
	fibonacciSelect(cc, quit)

	// fmt.Println("==== AFTER SELECT ====")
	// tick := time.Tick(100 * time.Millisecond)
	// boom := time.After(500 * time.Millisecond)
	// start := time.Now()
	// elapsed := func() time.Duration {
	// 	return time.Since(start).Round(time.Millisecond)
	// }
	// // The default case in a select is run if no other case is ready
	// // Use a default case to try a send or receive without blocking
	// for {
	// 	select {
	// 	case <-tick:
	// 		fmt.Printf("[%6s] tick.\n", elapsed())
	// 	case <-boom:
	// 		fmt.Printf("[%6s] BOOM!\n", elapsed())
	// 		return
	// 	default:
	// 		fmt.Printf("[%6s]     .\n", elapsed())
	// 		time.Sleep(25 * time.Millisecond)
	// 	}
	// }

	fmt.Println("==== MUTEX ====")
	sc := SafeCounter{v: make(map[string]int)}
	for range 1000 {
		go sc.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(sc.Value("somekey"))
}
