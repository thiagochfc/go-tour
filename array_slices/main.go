package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	fmt.Println("==== Array ====")
	BasicExampleArray()
	fmt.Println("")

	fmt.Println("==== Slice ====")
	BasicExampleSlice()
	ExampleSliceWithString()
	ExampleSliceLiterals()
	fmt.Println("")

	a := []int{2, 3, 5, 7, 11, 13}
	printSlice(a)

	b := make([]int, 5)
	printSlice(b)

	c := make([]int, 0, 5)
	printSlice(c)

	TicTacToeBoard()

	d := make([]int, 0, 5)
	d = append(d, 1)
	printSlice(d)
	d = append(d, 2, 3, 4)
	printSlice(d)

	// i = index
	// v = value
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
