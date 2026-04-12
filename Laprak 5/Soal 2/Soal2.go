package main

import (
	"fmt"
	"strings"
)

func cetakBintang(n int) {
	if n <= 0 {
		return
	}
	cetakBintang(n - 1)

	fmt.Println(strings.Repeat("*", n))
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	cetakBintang(n)
}