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

func galiKartu(deck *Dominoes, target Domino) {
	fmt.Printf("Mencari kecocokan suit untuk target [%d|%d]...\n", target.Sisi1, target.Sisi2)

	for deck.Tersisa > 0 {
		drawn := ambilKartu(deck)
		fmt.Printf("  Menggali kartu: [%d|%d] -> ", drawn.Sisi1, drawn.Sisi2)

		cocokSisi1 := gambarKartu(drawn, 1) == gambarKartu(target, 1) || gambarKartu(drawn, 1) == gambarKartu(target, 2)
		cocokSisi2 := gambarKartu(drawn, 2) == gambarKartu(target, 1) || gambarKartu(drawn, 2) == gambarKartu(target, 2)

		if cocokSisi1 || cocokSisi2 {
			fmt.Println("Cocok! Berhenti menggali.")
			return
		} else {
			fmt.Println("Tidak cocok.")
		}
	}

	fmt.Println("  Tumpukan habis. Tidak ada kartu yang cocok.")
}

func sepasangKartu(d1 Domino, d2 Domino) bool {
	total := nilaiKartu(d1) + nilaiKartu(d2)
	return total == 12
}

func main() {
	var deck Dominoes
	inisialisasiKartu(&deck)
	kocokKartu(&deck)

	fmt.Println("==== TEST SOAL 2.A: GALI KARTU ====")
	kartuTarget := Domino{Sisi1: 3, Sisi2: 5, Nilai: 8, Balak: false}
	galiKartu(&deck, kartuTarget)

	fmt.Println("\n==== TEST SOAL 2.B: SEPASANG KARTU ====")
	kartuA := Domino{Sisi1: 6, Sisi2: 6, Nilai: 12, Balak: true}
	kartuB := Domino{Sisi1: 0, Sisi2: 0, Nilai: 0, Balak: true}
	fmt.Printf("Kartu A [%d|%d] dan Kartu B [%d|%d]. Apakah total 12? %v\n", 
		kartuA.Sisi1, kartuA.Sisi2, kartuB.Sisi1, kartuB.Sisi2, sepasangKartu(kartuA, kartuB))

	kartuC := Domino{Sisi1: 3, Sisi2: 4, Nilai: 7, Balak: false}
	kartuD := Domino{Sisi1: 1, Sisi2: 4, Nilai: 5, Balak: false}
	fmt.Printf("Kartu C [%d|%d] dan Kartu D [%d|%d]. Apakah total 12? %v\n", 
		kartuC.Sisi1, kartuC.Sisi2, kartuD.Sisi1, kartuD.Sisi2, sepasangKartu(kartuC, kartuD))

	kartuE := Domino{Sisi1: 1, Sisi2: 2, Nilai: 3, Balak: false}
	fmt.Printf("Kartu A [%d|%d] dan Kartu E [%d|%d]. Apakah total 12? %v\n", 
		kartuA.Sisi1, kartuA.Sisi2, kartuE.Sisi1, kartuE.Sisi2, sepasangKartu(kartuA, kartuE))
}