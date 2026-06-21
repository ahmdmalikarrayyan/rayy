package main

import (
	"fmt"
)

func main() {
	var x string
	var n int

	fmt.Scan(&x)

	fmt.Scan(&n)

	data := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	var ditemukan bool = false
	var posisi []int
	var jumlah int = 0

	for i := 0; i < n; i++ {
		if data[i] == x {
			ditemukan = true
			posisi = append(posisi, i+1)
			jumlah++
		}
	}

	minimalDua := jumlah >= 2

	fmt.Println("========================================")
	
	if ditemukan {
		fmt.Println("a. Apakah string x ada? Ya, ada.")
	} else {
		fmt.Println("a. Apakah string x ada? Tidak ada.")
	}

	if ditemukan {
		fmt.Printf("b. Posisi string x ditemukan: %v\n", posisi)
	} else {
		fmt.Println("b. Posisi string x ditemukan: -")
	}

	fmt.Printf("c. Jumlah string x dalam kumpulan data: %d\n", jumlah)

	if minimalDua {
		fmt.Println("d. Adakah sedikitnya dua string x? Ya.")
	} else {
		fmt.Println("d. Adakah sedikitnya dua string x? Tidak.")
	}
	fmt.Println("========================================")
}