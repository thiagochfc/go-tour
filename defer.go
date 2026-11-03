package main

import "fmt"

func deferWithoutLoop() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func deferWithLoop() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	deferWithoutLoop()
	deferWithLoop()
}
