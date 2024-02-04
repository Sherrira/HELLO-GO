package main

import "fmt"

func soma(a, b *int) int {
	*a = 30
	*b = 1
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	x := soma(&minhaVar1, &minhaVar2)
	fmt.Println(x)

	fmt.Println(minhaVar1)
}