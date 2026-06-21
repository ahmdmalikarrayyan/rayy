package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	var tetesanA, tetesanB, tetesanC, tetesanD int

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x < 0.5 && y < 0.5 {
			tetesanA++
		} else if x >= 0.5 && y < 0.5 {
			tetesanB++
		} else if x >= 0.5 && y >= 0.5 {
			tetesanC++
		} else if x < 0.5 && y >= 0.5 {
			tetesanD++
		}
	}

	curahA := float64(tetesanA) * 0.0001
	curahB := float64(tetesanB) * 0.0001
	curahC := float64(tetesanC) * 0.0001
	curahD := float64(tetesanD) * 0.0001

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahD)
}