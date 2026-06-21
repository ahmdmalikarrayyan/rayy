package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	pi_prev := 0.0
	pi_curr := 0.0

	for i := 1; i <= n; i++ {
		term := 4.0 / float64(2*i-1)
		
		if i%2 == 0 {
			pi_curr = pi_prev - term
		} else {
			pi_curr = pi_prev + term
		}

		if i > 1 {
			selisih := math.Abs(pi_curr - pi_prev)
			
			if selisih <= 0.00001 {
				fmt.Printf("Hasil PI: %.10f\n", pi_prev)
				fmt.Printf("Hasil PI: %.10f\n", pi_curr)
				
				fmt.Printf("Pada i ke: %d\n", i+1) 
				break
			}
		}
		
		pi_prev = pi_curr
	}
}