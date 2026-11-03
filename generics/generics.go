package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if x == v {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next  *List[T]
	value T
}

func (l *List[T]) Add(v T) {
	l.next = &List[T]{nil, v}
}

func (l List[T]) Print() {
	ll := &l
	for {
		if ll == nil {
			break
		}

		fmt.Printf("Valor = %v\n", ll.value)
		ll = ll.next
	}
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	l := List[int]{nil, 1}
	l.Add(3)
	l.Print()
}
