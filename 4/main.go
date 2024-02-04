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
	f CPF = 36794311801
)

func main() {

	fmt.Println(b, c, d, e)
	fmt.Printf("Segue o número do CPF: %v ", f)
}
