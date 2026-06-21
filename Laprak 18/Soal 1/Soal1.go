package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	Sisi1 int
	Sisi2 int
	Nilai int
	Balak bool
}

type Dominoes struct {
	Kartu   [28]Domino
	Tersisa int
}

func kocokKartu(deck *Dominoes) {
	rand.Seed(time.Now().UnixNano())

	for i := deck.Tersisa - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck.Kartu[i], deck.Kartu[j] = deck.Kartu[j], deck.Kartu[i]
	}
}

func ambilKartu(deck *Dominoes) Domino {
	if deck.Tersisa == 0 {
		return Domino{}
	}
	deck.Tersisa--
	return deck.Kartu[deck.Tersisa]
}

func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.Sisi1
	} else if suit == 2 {
		return d.Sisi2
	}
	return -1
}

func nilaiKartu(d Domino) int {
	return d.Nilai
}


func inisialisasiKartu(deck *Dominoes) {
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			deck.Kartu[idx] = Domino{
				Sisi1: i,
				Sisi2: j,
				Nilai: i + j,
				Balak: i == j,
			}
			idx++
		}
	}
	deck.Tersisa = 28
}

func main() {
	var setDomino Dominoes

	inisialisasiKartu(&setDomino)
	fmt.Println("Jumlah kartu awal:", setDomino.Tersisa)

	kocokKartu(&setDomino)
	fmt.Println("Kartu telah dikocok...")

	kartuDiTangan := ambilKartu(&setDomino)
	fmt.Println("\n-- MENGAMBIL KARTU --")
	fmt.Printf("Kartu ditarik: [%d|%d]\n", kartuDiTangan.Sisi1, kartuDiTangan.Sisi2)
	fmt.Println("Sisa kartu di deck:", setDomino.Tersisa)

	fmt.Println("\n-- CEK GAMBAR (SUIT) --")
	fmt.Println("Sisi pertama:", gambarKartu(kartuDiTangan, 1))
	fmt.Println("Sisi kedua  :", gambarKartu(kartuDiTangan, 2))

	fmt.Println("\n-- CEK NILAI DAN STATUS --")
	fmt.Println("Nilai Total :", nilaiKartu(kartuDiTangan))
	
	if kartuDiTangan.Balak {
		fmt.Println("Status      : Ini adalah kartu BALAK (Kembar)!")
	} else {
		fmt.Println("Status      : Bukan kartu balak.")
	}
}