package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	var beratIkan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := int(math.Ceil(float64(x) / float64(y)))
	
	totalBeratWadah := make([]float64, jumlahWadah)
	var totalSemuaWadah float64

	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalBeratWadah[indeksWadah] += beratIkan[i]
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", totalBeratWadah[i])
		totalSemuaWadah += totalBeratWadah[i]
	}
	fmt.Println()

	rataRata := totalSemuaWadah / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rataRata)
}