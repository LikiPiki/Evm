package main

import (
	"fmt"
)

type InputReader struct {
	OperationCount int
}

func NewReader() InputReader {
	return InputReader{}
}

func (r *InputReader) Read() {
	fmt.Print("Enter command count: ")
	fmt.Scanf("%d", &r.OperationCount)
	fmt.Println(r)
}
