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

		if num == -5313 {
			break
		}

		if num == 0 {
			if len(data) > 0 {
				insertionSort(data)

				n := len(data)
				var median int

				if n%2 != 0 {
					median = data[n/2]
				} else {
					mid1 := data[(n/2)-1]
					mid2 := data[n/2]
					median = (mid1 + mid2) / 2
				}

				fmt.Println(median)
			}
		} else {
			data = append(data, num)
		}
	}
}