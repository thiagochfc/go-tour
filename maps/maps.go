package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

const (
	BellLabs = "Bell Labs"
	Google   = "Google"
	Answer   = "Answer"
	TheValue = "The value:"
	EmptyStr = " "
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	strSplit := strings.SplitSeq(s, EmptyStr)

	for str := range strSplit {
		v, ok := result[str]
		if !ok {
			result[str] = 0
		}
		result[str] = v + 1
	}

	return result
}

func main() {
	m := make(map[string]Vertex)
	m[BellLabs] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m[BellLabs])

	var n = map[string]Vertex{
		BellLabs: {
			40.68433, -74.39967,
		},
		Google: {
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)

	o := make(map[string]int)

	o[Answer] = 42
	fmt.Println(TheValue, o[Answer])

	o[Answer] = 48
	fmt.Println(TheValue, o[Answer])

	delete(o, Answer)
	fmt.Println(TheValue, m["Answer"])

	v, ok := m["Answer"]
	fmt.Println(TheValue, v, "Present?", ok)

	wc.Test(WordCount)
}
