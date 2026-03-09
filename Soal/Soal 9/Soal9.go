package main

import (
	"fmt"
)

func main() {
	var K int

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	akar2 := 1.0

	for k := 0; k <= K; k++ {
		fk := float64(k)
		
		pembilang := (4*fk + 2) * (4*fk + 2)
		penyebut := (4*fk + 1) * (4*fk + 3)
		
		akar2 *= (pembilang / penyebut)
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}