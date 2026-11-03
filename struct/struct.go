package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Vortex struct {
	X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	vertex := Vertex{3, 4}
	vertex.X = 5
	fmt.Println(vertex, vertex.Y)

	vortex := Vortex{7, 8}
	v := &vortex
	v.X = 1e9
	fmt.Println(*v)
}
