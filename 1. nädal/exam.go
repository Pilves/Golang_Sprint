package main

import "fmt"

// ChessBoard funktsioon genereerib malelaua ruutude positsioonid vastavalt värvile.
func ChessBoard(black bool) string {
    var blacks string // Hoiab mustade ruutude positsioone.
    var whites string // Hoiab valgete ruutude positsioone.
    // Malelaua ruudud on määratletud kahe tsükliga: üks jookseb ridade ja teine veergude läbi.
    for j := 1; j <= 8; j++ {
        for i := 'a'; i <= 'h'; i++ {
            // Kontrollib, kas veeru ja rea indeksite summa on paarisarv.
            if (int(i)+j) % 2 == 0 {
                // Lisab mustade ruutude nimekirja.
                if blacks == "" {
                    blacks = string(i) + string(j+'0')
                } else {
                    blacks += ", " + string(i) + string(j+'0')
                }
            } else {
                // Lisab valgete ruutude nimekirja.
                if whites == "" {
                    whites = string(i) + string(j+'0')
                } else {
                    whites += ", " + string(i) + string(j+'0')
                }
            }
        }
    }
    // Tagastab kas mustade või valgete ruutude positsioonid sõltuvalt funktsiooni sisendist.
    if black {
        return blacks
    } else {
        return whites
    }
}

// Interest funktsioon arvutab, mitu aastat kulub algsumma kasvatamiseks lõppsummaks aastase intressimääraga.
func Interest(start, finish, interest float64) int {
    years := 0 // Aastate loendur.
    if start >= finish {
        return 0 // Kui algsumma on juba suurem või võrdne lõppsummaga, tagasta 0.
    }
    for i := start; i < finish; i +=(i*interest) {
        years++
        if years >= 99 {
            return 0 // Tagastab 0, kui aastaid koguneb 99 või rohkem.
        }
    }
    return years // Tagastab kogunenud aastate arvu.
}

// Optimeeritud interest funktsioon arvutab effektiivsemalt, mitu aastat kulub start-il finishini jõudmiseks aastase intressimääraga.
func InterestOptimized(start, finish, interest float64) int {
    // Kontrollib, kas algsumma on juba suurem või võrdne lõppsummaga.
    if start >= finish {
        return 0 // Kui jah, siis funktsioon lõpetab töö ja tagastab 0.
    }

    years := 0 // Loob muutuja aastate loendamiseks, alustades väärtusega 0.
    // Tsükkel käivitatakse algsummast ja jätkub, kuni summa on väiksem kui lõppsumma.
    for i := start; i < finish; years++ {
        // Kui aastate arv jõuab 99-ni, lõpetab funktsioon töö ja tagastab 0.
        if years >= 99 {
            return 0
        }
        i += i * interest // Suurendab investeeringu summat, lisades sellele intressi osa.
    }
    return years // Tagastab kogunenud aastate arvu, mis kulus algsumma kasvamiseks lõppsummani.
}

func main(){
    fmt.Println(ChessBoard(true)) // Prindib mustade ruutude positsioonid malelaual. || OUTPUT: "a1, c1, e1, g1, b2....b4, d4, f4, h4, a5, c5....b8, d8, f8, h8"
    fmt.Println(Interest(100.0, 121.0, 0.1)) // Prindib aastate arvu, mis kulub 100 kasvatamiseks 121-ni 10% intressiga. || OUTPUT: 2
    fmt.Println(InterestOptimized(100.0, 121.0, 0.1)) // Prindib aastate arvu, mis kulub 100 kasvatamiseks 121-ni 10% intressiga. || OUTPUT: 2
}
