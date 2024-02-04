package main

import "fmt"


func main() {
	var minhaVar interface{} = "Sherr"
	fmt.Println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v.\n", res, ok)
	}


