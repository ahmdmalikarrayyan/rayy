package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nmrFact int
	
	factorial(n, &nFact)
	factorial(n-r, &nmrFact)
	
	*hasil = nFact / nmrFact
}

func combination(n, r int, hasil *int) {
	var nFact, rFact, nmrFact int
	
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &nmrFact)
	
	*hasil = nFact / (rFact * nmrFact)
}

func main() {
	var a, b, c, d int
	var resP, resC int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &resP)
	combination(a, c, &resC)
	fmt.Printf("%d %d\n", resP, resC)

	permutation(b, d, &resP)
	combination(b, d, &resC)
	fmt.Printf("%d %d\n", resP, resC)
}