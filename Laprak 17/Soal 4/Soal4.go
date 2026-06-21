package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	S := 0.0

	for i := 1; i <= n; i++ {
		term := 1.0 / float64(2*i-1)
		
		if i%2 == 0 {
			S -= term
		} else {
			S += term
		}
	}

	pi := S * 4.0
	fmt.Printf("Hasil PI: %.7f\n", pi)
}