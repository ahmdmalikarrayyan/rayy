package main

import (
	"fmt"
)

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}

	maxIdx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	
	fmt.Println(pustaka[maxIdx].judul, pustaka[maxIdx].penulis, pustaka[maxIdx].penerbit, pustaka[maxIdx].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1

		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}

	for i := 0; i < limit; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(pustaka[i].judul)
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	left := 0
	right := n - 1
	foundIdx := -1

	for left <= right {
		mid := (left + right) / 2

		if pustaka[mid].rating == r {
			foundIdx = mid
			break
		} else if r > pustaka[mid].rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]
		fmt.Println(b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var Pustaka DaftarBuku
	var nPustaka int

	DaftarkanBuku(&Pustaka, &nPustaka)

	CetakTerfavorit(Pustaka, nPustaka)

	UrutBuku(&Pustaka, nPustaka)

	Cetak5Terbaru(Pustaka, nPustaka)

	var ratingDicari int
	_, err := fmt.Scan(&ratingDicari)
	if err == nil {
		CariBuku(Pustaka, nPustaka, ratingDicari)
	}
}