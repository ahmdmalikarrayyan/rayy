package main

import (
	"fmt"
)

func main() {
	var bilangan float64
	var total float64 = 0
	var jumlahData int = 0

	fmt.Println("Masukkan bilangan real (ketik 9999 untuk berhenti dan melihat hasil):")

	for {
		_, err := fmt.Scan(&bilangan)
		if err != nil {
			break
		}

		if bilangan == 9999 {
			break
		}

		total += bilangan
		jumlahData++
	}

	if jumlahData > 0 {
		rerata := total / float64(jumlahData)
		fmt.Printf("Rerata dari bilangan tersebut adalah: %.2f\n", rerata)
	} else {
		fmt.Println("Tidak ada bilangan yang dimasukkan untuk dihitung reratanya.")
	}
}