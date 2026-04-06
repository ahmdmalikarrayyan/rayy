package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var totalSoal, totalSkor int
	var maxSoal, minSkor int

	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&totalSoal, &totalSkor)

		if totalSoal > maxSoal {
			maxSoal = totalSoal
			minSkor = totalSkor
			pemenang = nama
		} else if totalSoal == maxSoal {
			if totalSkor < minSkor {
				minSkor = totalSkor
				pemenang = nama
			}
		}
	}

	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
	}
}