package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for x := range b {
		b[x] = 'A'
	}
	return len(b), nil
}

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)

		if err == io.EOF {
			break
		}

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
	}

	reader.Validate(MyReader{})
}
