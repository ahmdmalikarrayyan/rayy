package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 0 {
		fmt.Println("Input harus lebih besar dari 0")
		return
	}

	fmt.Print("Faktor: ")
	jumlahFaktor := 0

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++
		}
	}
	fmt.Println()

	fmt.Print("Prima: ")
	switch jumlahFaktor {
	case 2:
		fmt.Println("true")
	default:
		fmt.Println("false")
	}
}