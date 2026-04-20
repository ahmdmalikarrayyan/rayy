package main

import (
	"fmt"
	"math"
)

const MAX_CAPACITY int = 100

func main() {
	var arr [MAX_CAPACITY]int
	var n, x, idx, target int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan %d elemen:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("\na. Isi array: ")
	tampilkanSemua(arr, n)

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("d. Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("   Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	avg := hitungRataRata(arr, n)
	fmt.Printf("f. Rata-rata: %.2f\n", avg)

	sd := hitungStandarDeviasi(arr, n, avg)
	fmt.Printf("g. Standar Deviasi: %.2f\n", sd)

	fmt.Print("h. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&target)
	fmt.Printf("   Frekuensi %d: %d kali\n", target, hitungFrekuensi(arr, n, target))

	fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	n = hapusElemen(&arr, n, idx)
	fmt.Print("   Isi array setelah dihapus: ")
	tampilkanSemua(arr, n)
}

func tampilkanSemua(arr [MAX_CAPACITY]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func hitungRataRata(arr [MAX_CAPACITY]int, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n)
}

func hitungStandarDeviasi(arr [MAX_CAPACITY]int, n int, avg float64) float64 {
	var varians float64
	for i := 0; i < n; i++ {
		varians += math.Pow(float64(arr[i])-avg, 2)
	}
	return math.Sqrt(varians / float64(n))
}

func hitungFrekuensi(arr [MAX_CAPACITY]int, n int, target int) int {
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == target {
			count++
		}
	}
	return count
}

func hapusElemen(arr *[MAX_CAPACITY]int, n int, idx int) int {
	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	return n - 1 
}