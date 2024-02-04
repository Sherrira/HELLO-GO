package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(2, 7)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum (a, b int) (int, error) {
	if a*b>= 65{ 
	return 0, errors.New("Sono")
	}
	return a*b, nil
}