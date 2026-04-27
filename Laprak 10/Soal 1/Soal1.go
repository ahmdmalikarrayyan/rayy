package main

import (
	"fmt"
)

func main() {
	var beratKelinci [1000]float64
	var n int

	fmt.Scan(&n)

	if n > 1000 {
		n = 1000
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	if n > 0 {
		min := beratKelinci[0]
		max := beratKelinci[0]

		for i := 1; i < n; i++ {
			if beratKelinci[i] < min {
				min = beratKelinci[i]
			}
			if beratKelinci[i] > max {
				max = beratKelinci[i]
			}
		}

		fmt.Printf("%.2f %.2f\n", min, max)
	}
}