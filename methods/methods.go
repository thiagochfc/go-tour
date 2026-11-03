package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) AbsPointer() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Choosing a value or pointer receiver
// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
// Continue...
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	v.Scale(10)
	fmt.Println(v.Abs())

	v = Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))

	v = Vertex{3, 4}
	fmt.Println((&v).Abs())

	x := &v
	fmt.Println((*x).Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Continuing...
	// In this example, both Scale and AbsPointer are methods with receiver type *Vertex, even though the Abs method needn't modify its receiver.
	// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
	z := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", z, z.AbsPointer())
	z.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", z, z.AbsPointer())
}
