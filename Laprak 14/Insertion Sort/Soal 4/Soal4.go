package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var data []int
	var num int

	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		
		if num < 0 {
			break
		}
		
		data = append(data, num)
	}

	if len(data) > 0 {
		insertionSort(data)

		for i, val := range data {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()

		if len(data) == 1 {
			fmt.Println("Data berjarak 0")
		} else {
			jarakTetap := true
			jarak := data[1] - data[0]

			for i := 2; i < len(data); i++ {
				if data[i]-data[i-1] != jarak {
					jarakTetap = false
					break
				}
			}

			if jarakTetap {
				fmt.Printf("Data berjarak %d\n", jarak)
			} else {
				fmt.Println("Data berjarak tidak tetap")
			}
		}
	}
}