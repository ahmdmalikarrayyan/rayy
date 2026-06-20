package main

import (
	"fmt"
)

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var num int
			fmt.Scan(&num)
			
			if num%2 != 0 {
				ganjil = append(ganjil, num)
			} else {
				genap = append(genap, num)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		first := true
		
		for _, val := range ganjil {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(val)
			first = false
		}
		
		for _, val := range genap {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(val)
			first = false
		}
		
		fmt.Println()
	}
}