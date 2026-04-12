package main

import "fmt"

func fibboacci(n int) int{
	if n <= 1 {
		return n
	}
	return fibboacci(n-1) + fibboacci(n-2)
}

func main() {
	fmt.Println("Deret Fibboacci hingga suku ke-10")
	fmt.Println("------------------------------")
	fmt.Printf("| %-2s | %-3s |\n", "n", "Sn")
	fmt.Println("------------------------------")

	for i := 0; i <= 10; i++ {
		result := fibboacci(i)
		fmt.Printf("| %-2d | %-3d |\n", i, result)
	}
	fmt.Println("------------------------------")
}