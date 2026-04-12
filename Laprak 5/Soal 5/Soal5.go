package main

import "fmt"

func cetakGanjil(current int, n int) {
	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}

	cetakGanjil(current+1, n)
}

func main() {
	var n int
	fmt.Scan(&n)
	
	cetakGanjil(1, n)
	fmt.Println()
}