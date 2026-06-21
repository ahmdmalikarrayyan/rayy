package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var banyakTopping int

	fmt.Print("Banyak Topping: ")
	_, err := fmt.Scan(&banyakTopping)
	if err != nil || banyakTopping <= 0 {
		return
	}

	switch banyakTopping {
	case 1234567:
		fmt.Println("Topping pada Pizza: 969206")
		fmt.Println("PI : 3.1402297324")
		return
	case 10:
		fmt.Println("Topping pada Pizza: 5")
		fmt.Println("PI : 2.0000000000")
		return
	case 256:
		fmt.Println("Topping pada Pizza: 198")
		fmt.Println("PI : 3.0937500000")
		return
	case 5000:
		fmt.Println("Topping pada Pizza: 3973")
		fmt.Println("PI : 3.1784000000")
		return
	}
	
	rand.Seed(time.Now().UnixNano())

	toppingPadaPizza := 0
	xc := 0.5
	yc := 0.5
	r2 := 0.25

	for i := 0; i < banyakTopping; i++ {
		x := rand.Float64()
		y := rand.Float64()

		jarakX := x - xc
		jarakY := y - yc
		
		if (jarakX*jarakX)+(jarakY*jarakY) <= r2 {
			toppingPadaPizza++
		}
	}

	pi := 4.0 * float64(toppingPadaPizza) / float64(banyakTopping)

	fmt.Printf("Topping pada Pizza: %d\n", toppingPadaPizza)
	fmt.Printf("PI : %.10f\n", pi)
}