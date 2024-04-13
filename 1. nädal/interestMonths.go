package main

import (
	"fmt"
	"math"
)

// Funktsioon arvutab eesmärgini jõudmise aastate ja kuude kaupa
func InterestWithMonths(start, finish, interest float64) (int, int) {
	if start >= finish {     // Kui algus on suurem või sama lõppeesmärgiga, lõpeta kohe
		return 0, 0
	}
	years := 0              // Aastate muutuja
	months := 0             // Kuude muutuja
	monthlyInterest := math.Pow(1+interest, 1.0/12) - 1  // Arvuta kuupõhine kasvumäär (month based compounding interest aka kuupõhine liitintress)

	for start < finish {    //kuniks algus on väiksem kui lõpp
		start += start * monthlyInterest  // Kasvata algust iga kuu intressi võrra
		months++                          // Lisa üks kuu
		if start >= finish {              // Kui oleme lõpus, vaata, mitu kuud on läinud
			if months >= 12 {             // Kui kuid on 12 või rohkem
				years += months / 12     // Arvuta uued aastad
				months = months % 12     // Jäta allesjäänud kuud
			}
			return years, months         // Tagasta aastad ja kuud
		}
	}
	return years, months                // Tagasta aastad ja kuud (siia kohta ei tohiks kood kunagi jõuda)
}


func main() {
	years, months := InterestWithMonths(6540, 1210000, 0.4)
	fmt.Printf("eesmärgini jõudmiseks kulub %d aasta(t) ja %d kuu(d)\n", years, months)
}
