package main

/*
Basic types

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func addInt(x int, y int) int {
	return x + y
}

func add2Float(x, y float64) float64 {
	return x + y
}

func swap(first, second string) (string, string) {
	return second, first
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

var z, u int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	t      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("The time is: ", time.Now())

	fmt.Println("My number favorite is", rand.Intn(10))

	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))

	fmt.Println(addInt(42, 13))

	fmt.Println(add2Float(33.33, 66.77))

	fmt.Println(swap("Hello", "World"))

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	var news, sames, olds = true, false, "no!"
	fmt.Println(z, u, news, sames, olds)

	k := 3
	fmt.Println(k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", t, t)

	var o, p int = 3, 4
	var f float64 = math.Sqrt(float64(o*o + p*p))
	var b uint = uint(f)
	fmt.Println(o, p, b)

	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
