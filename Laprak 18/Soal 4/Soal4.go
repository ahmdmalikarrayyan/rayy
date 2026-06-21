package main

import (
	"fmt"
)


type MesinKarakter struct {
	pita   string
	indeks int
}

func (m *MesinKarakter) start(teks string) {
	m.pita = teks
	m.indeks = 0
}

func (m *MesinKarakter) maju() {
	m.indeks++
}

func (m *MesinKarakter) eop() bool {
	if m.indeks >= len(m.pita) {
		return true
	}
	return m.pita[m.indeks] == '.'
}

func (m *MesinKarakter) cc() byte {
	return m.pita[m.indeks]
}


func main() {
	var mesin MesinKarakter

	teksInput := "ALAMAT LENGKAP JALAN LEMBONG NO. LIMA."
	
	mesin.start(teksInput)

	var totalKarakter int = 0
	var totalA int = 0
	var totalLE int = 0
	
	var charSebelumnya byte = ' ' 

	fmt.Print("Teks terbaca: ")

	for !mesin.eop() {
		karakterSaatIni := mesin.cc()
		
		fmt.Print(string(karakterSaatIni))

		totalKarakter++

		if karakterSaatIni == 'A' {
			totalA++
		}

		if charSebelumnya == 'L' && karakterSaatIni == 'E' {
			totalLE++
		}

		charSebelumnya = karakterSaatIni

		mesin.maju()
	}

	fmt.Println()
	fmt.Println("========================================")

	var frekuensiA float64 = 0.0
	if totalKarakter > 0 {
		frekuensiA = float64(totalA) / float64(totalKarakter)
	}

	fmt.Printf("1. Total Karakter Terbaca   : %d\n", totalKarakter)
	fmt.Printf("2. Total Huruf 'A'          : %d\n", totalA)
	fmt.Printf("3. Frekuensi Huruf 'A'      : %.4f (atau %.2f%%)\n", frekuensiA, frekuensiA*100)
	fmt.Printf("4. Total Pasangan kata 'LE' : %d\n", totalLE)
}