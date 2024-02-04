package main

import "fmt"

type MyNumber int
type Number interface{
	~int | float64
}

func Soma[T Number] (m map[string]T) T{
	var soma T
	for _, v := range m{
		soma += v		 
	}
	return soma
}

func Compara[T Number] (a T, b T) bool{
	if a==b{
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Sherr": 12000, "Fill":24000, "Irani": 5000}
	fmt.Println(Soma(m))
	m2 := map[string]float64{"Sherr": 12000.00, "Fill":24000.90, "Irani": 5000.25}
	fmt.Println(Soma(m2))
	m3 := map[string]MyNumber{"Sherr": 12000, "Fill":24000, "Irani": 5000}
	fmt.Println(Soma(m3))

	fmt.Println(Compara(10, 10.0))
	}


