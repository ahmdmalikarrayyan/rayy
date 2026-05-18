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

	max1 := -1
	max2 := -1
	ketua := 0
	wakil := 0

	for i := 1; i <= 20; i++ {
		suara := suaraKandidat[i]
		
		if suara > max1 {
			max2 = max1
			wakil = ketua
			
			max1 = suara
			ketua = i
		} else if suara > max2 {
			max2 = suara
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}