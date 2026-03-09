package main

import "fmt"

func main() {
	var celsius float64
	var fahrenheit, reamur, kelvin float64

	fmt.Print("Input Celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit = (9.0/5.0)*celsius + 32
	reamur = (4.0/5.0)*celsius
	kelvin = celsius + 273.15

	fmt.Printf("Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Reamur: %.2f\n", reamur)
	fmt.Printf("Kelvin: %.2f\n", kelvin)
}