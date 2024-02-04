package main

import ("fmt")


func main() {
	s := []int {2, 4, 8, 16, 32, 64, 128, 258, 512, 1024, 2048}

	fmt.Printf("len=%d - cap=%d - v=%v\n", len(s), cap(s), s)
	fmt.Printf("len=%d - cap=%d - v=%v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d - cap=%d - v=%v\n", len(s[:5]), cap(s[:5]), s[:5])
	fmt.Printf("len=%d - cap=%d - v=%v\n", len(s[3:7]), cap(s[3:7]), s[3:7])

	s=append (s,4096)
	fmt.Printf("len=%d - cap=%d - v=%v\n", len(s), cap(s), s)

	for i, v := range s [7:10]{
		fmt.Printf("%d - %d\n", i, v)
	}
}

