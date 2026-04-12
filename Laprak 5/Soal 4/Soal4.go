package main

import "fmt"

func tampilkanBarisan(n int) {
	if n == 1 {
		fmt.Print(1, " ")
		return
	}

	fmt.Print(n, " ")

	tampilkanBarisan(n - 1)

	fmt.Print(n, " ")
}

func main() {
	var input int
	fmt.Scan(&input)

	if input > 0 {
		tampilkanBarisan(input)
		fmt.Println()
	}
}