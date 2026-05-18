package main

import (
	"fmt"
)

func main() {
	var suaraKandidat [21]int

	var totalSuaraMasuk int = 0
	var totalSuaraSah int = 0
	var pilihan int

	for {
		_, err := fmt.Scan(&pilihan)
		
		if err != nil {
			break
		}

		if pilihan == 0 {
			break
		}

		totalSuaraMasuk++

		if pilihan >= 1 && pilihan <= 20 {
			totalSuaraSah++
			suaraKandidat[pilihan]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)

	for i := 1; i <= 20; i++ {
		if suaraKandidat[i] > 0 {
			fmt.Printf("%d : %d\n", i, suaraKandidat[i])
		}
	}
}