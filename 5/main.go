package main

import (
	"fmt"
)

const a = "GOGOGO"

type CPF int

var (
	b bool    = true
	c int     = 22
	d string  = "te amo"
	e float64 = 3.14
	f CPF     = 36794311801
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 22
	meuArray[2] = 300000

	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor no índice %d é %d\n", i, v)

	}

}
