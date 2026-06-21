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
}

type Dominoes struct {
	Kartu   [28]Domino
	Tersisa int
}

type Pemain struct {
	Nama  string
	Kartu []Domino
}

func inisialisasiKartu(deck *Dominoes) {
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			deck.Kartu[idx] = Domino{Sisi1: i, Sisi2: j, Nilai: i + j}
			idx++
		}
	}
	deck.Tersisa = 28
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

func main() {
	var deck Dominoes
	inisialisasiKartu(&deck)
	kocokKartu(&deck)

	pemain := []Pemain{
		{Nama: "Pemain 1"},
		{Nama: "Pemain 2"},
		{Nama: "Pemain 3"},
		{Nama: "Pemain 4"},
	}

	for i := 0; i < 4; i++ {
		for k := 0; k < 7; k++ {
			pemain[i].Kartu = append(pemain[i].Kartu, ambilKartu(&deck))
		}
	}

	fmt.Println("=== MEMULAI SIMULASI GAPLEH ===")
	
	ujungKiri := -1
	ujungKanan := -1
	
	giliran := 0
	passBerturut := 0

	for {
		p := &pemain[giliran]
		berhasilJalan := false

		if ujungKiri == -1 && ujungKanan == -1 {
			dimainkan := p.Kartu[0]
			ujungKiri = dimainkan.Sisi1
			ujungKanan = dimainkan.Sisi2
			
			p.Kartu = p.Kartu[1:] 
			berhasilJalan = true
			fmt.Printf("%s memulai game dengan mengeluarkan [%d|%d]\n", p.Nama, dimainkan.Sisi1, dimainkan.Sisi2)
		} else {
			for i, k := range p.Kartu {
				if k.Sisi1 == ujungKiri {
					ujungKiri = k.Sisi2
					berhasilJalan = true
					fmt.Printf("%s menyambung di kiri dengan [%d|%d]. (Ujung Meja: Kiri %d, Kanan %d)\n", p.Nama, k.Sisi1, k.Sisi2, ujungKiri, ujungKanan)
				} else if k.Sisi2 == ujungKiri {
					ujungKiri = k.Sisi1
					berhasilJalan = true
					fmt.Printf("%s menyambung di kiri dengan [%d|%d]. (Ujung Meja: Kiri %d, Kanan %d)\n", p.Nama, k.Sisi1, k.Sisi2, ujungKiri, ujungKanan)
				} else if k.Sisi1 == ujungKanan {
					ujungKanan = k.Sisi2
					berhasilJalan = true
					fmt.Printf("%s menyambung di kanan dengan [%d|%d]. (Ujung Meja: Kiri %d, Kanan %d)\n", p.Nama, k.Sisi1, k.Sisi2, ujungKiri, ujungKanan)
				} else if k.Sisi2 == ujungKanan {
					ujungKanan = k.Sisi1
					berhasilJalan = true
					fmt.Printf("%s menyambung di kanan dengan [%d|%d]. (Ujung Meja: Kiri %d, Kanan %d)\n", p.Nama, k.Sisi1, k.Sisi2, ujungKiri, ujungKanan)
				}

				if berhasilJalan {
					p.Kartu = append(p.Kartu[:i], p.Kartu[i+1:]...)
					break
				}
			}
		}

		if berhasilJalan {
			passBerturut = 0
			
			if len(p.Kartu) == 0 {
				fmt.Printf("\n🎉 GAME OVER! %s MENANG karena kartunya habis!\n", p.Nama)
				break
			}
		} else {
			fmt.Printf("%s PASS (Tidak punya kartu yang cocok)\n", p.Nama)
			passBerturut++
			
			if passBerturut == 4 {
				fmt.Println("\n🛑 GAME OVER! STATUS GAPLEH (Semua pemain buntu).")
				hitungSkorAkhir(pemain)
				break
			}
		}

		giliran = (giliran + 1) % 4
	}
}

func hitungSkorAkhir(pemain []Pemain) {
	fmt.Println("--- Menghitung Sisa Poin ---")
	for _, p := range pemain {
		totalPoin := 0
		for _, k := range p.Kartu {
			totalPoin += k.Nilai
		}
		fmt.Printf("%s sisa %d kartu (Total Poin: %d)\n", p.Nama, len(p.Kartu), totalPoin)
	}
	fmt.Println("*Catatan: Pemain dengan poin terkecil memenangkan Gapleh.")
}