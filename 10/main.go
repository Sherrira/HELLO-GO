package main

import (
	"fmt"
)

func main() {
	total := func()int{
		return sum (2,4,6)*2
	}()

	fmt.Println(total)

}
func sum(numeros... int)int{

	total := 0
	for _, num := range numeros {
		total += num
	}
	return total

}