package main

import ("fmt")


func main() {
	sal := map[string]int{}
	sal ["Shé"] = 4800 
	sal ["Fill"] = 5900
	sal ["Iga"] = 3500
	sal ["Emo"] = 3700

	for _, sal := range sal {
		fmt.Printf("O salário é %d\n", sal)
		
	}


}