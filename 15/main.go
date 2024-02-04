package main

import "fmt"

func main() {
	a := 10
	var ponteiro *int = &a
	*ponteiro = 34
	b := &a
	fmt.Println(*b)
}