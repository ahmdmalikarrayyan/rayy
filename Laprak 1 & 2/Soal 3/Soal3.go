package main

import "fmt"

func main() {
	var r int
	var volume, luas float64
	pi := 3.1415926535

	fmt.Print("Jari-jari = ")
	fmt.Scanln(&r)

	volume = (4.0 / 3.0) * pi * float64(r*r*r)
	luas = 4 * pi * float64(r*r)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}