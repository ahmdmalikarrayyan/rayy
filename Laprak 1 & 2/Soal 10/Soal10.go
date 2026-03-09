package main

import "fmt"

func main() {
	var beratGram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	switch {
	case sisaGram >= 500:
		biayaSisa = sisaGram * 5
	case sisaGram < 500 && sisaGram > 0:
		biayaSisa = sisaGram * 15
	}

	if kg >= 10 {
		biayaSisa = 0
	}

	totalBiaya := biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}