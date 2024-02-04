package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(191,29,9,8))
}

func sum(numeros... int)int{

	total := 0
	for _, num := range numeros {
		total += num
	}
	return total

}