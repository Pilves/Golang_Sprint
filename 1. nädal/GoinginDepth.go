package main // See rida ütleb, et meie kood kuulub "main" paketti, mis on peamine programmipunkt.

import ( // "import" ütleb, et me kasutame väliseid packageid või libareyd.
	"fmt" // Kasutame fmt libaryt, et väljastada teksti või muutujaid.
)

/*
Alphabet Mastery
	Instructions
Create a function that takes a positive integer as input and returns the corresponding number of letters from the Latin alphabet. The input integer won't be larger than the alphabet's length.
*/
func AlphabetMastery(n int) string { // Defineerib funktsiooni nimega AlphabetMastery, mis võtab ühe täisarvu ja tagastab stringi.
	var answer string // Loob tühja stringi muutuja "answer" vastuste hoidmiseks.
	for i := 'a'; i < rune(n+'a'); i++ { // Käivitab tsükli, mis alustab tähest 'a' ja lõppeb enne n-ndat tähte alates 'a'-st.
		answer += string(i) // Lisab igal tsükli sammul järgmise tähe muutuja "answer" lõppu.
	}
	return answer // Tagastab kogutud tähtede stringi.
}


/*
Reverse Alphabet
	Instructions
Create a function that takes a step value n as input. Starting from z in the Latin alphabet, it should return the lowercase alphabet in reverse order as a string. For each letter, you skip n-1 letters. If n is 0 or negative, use a default step of 1.
*/
func ReverseAlphabet(step int) string {
	var answer string // Vastuse jaoks mõeldud string.
	if step <= 0 {    // Kui samm on 0 või negatiivne,
		step = 1 // kasutatakse vaikimisi sammu, mis on 1.
	}

	for i := 'z'; i >= 'a'; i -= rune(step) { // Käib läbi tähestiku 'z'-st 'a'-ni, hüpates üle n-1 tähe.
		answer += string(i) // Lisab iga tähe vastuse stringi.
	}
	return answer // Tagastab tagurpidi tähestiku.
}

/*
Pairs
	Instructions
Create a Go function that finds all possible combinations of two two-digit numbers. The pairs should be in ascending order, and each number should be padded with a leading zero if it's less than 10. The pairs should be separated by a comma and a space. Each number within a pair should be separated by a space. The function should return a string containing these pairs.
	For this task you're allowed to use:
		fmt
*/
func Pairs() string {
	var answer string          // Muutuja vastuse jaoks.
	for i := 0; i < 100; i++ { // Välise tsükli jaoks, mis läbib numbreid 0-st 99-ni.
		for j := i + 1; j < 100; j++ { // Sisemine tsükkel, loob paarid i-ga, alustades i+1-st.
			pair := fmt.Sprintf("%02d %02d", i, j) // Loob paari kahe numbri vahel, formaadis "xx xx".
			if answer == "" {                      // Kui vastus on tühi,
				answer = pair // määrab esimese paari vastuseks.
			} else { // Muul juhul
				answer += ", " + pair // lisab paarid vastusele, eraldatuna komaga.
			}
		}
	}
	return answer // Tagastab kõik paarid stringina.
}


/*
Combinations
	Instructions
Create a Go function that generates a series of unique triplets of digits. Each triplet should consist of different digits, and they should be in ascending order, from the smallest to the largest. The triplets should be separated by a comma and a space.
For example, the function should return triplets like 012, 013, 014, 015.... Avoid combinations with all the same digits, like 000 or 999, and exclude triplets in descending order, like 987.
	For this task you`re allowed to use:
		fmt
*/
func Combinations() string {
	var answer string         // Vastuse stringi algväärtustamine.
	for i := 0; i < 10; i++ { // Esimene tsükkel numbritele 0 kuni 9.
		for j := i + 1; j < 10; j++ { // Teine tsükkel, mis algab i-st +1, tagades et j on suurem kui i.
			for k := j + 1; k < 10; k++ { // Kolmas tsükkel, mis algab j-st +1, tagades et k on suurem kui j.
				triplet := fmt.Sprintf("%d%d%d", i, j, k) // Formeerib kolmekohalise numbri. XXX || (IJK)
				if answer == "" {                         // Kui vastus on tühi,
					answer = triplet // määrab answeri esimeseks kolmikuks.
				} else { // Muul juhul
					answer += ", " + triplet // Lisab kolmikud vastusele, eraldatuna komaga.
				}
			}
		}
	}
	return answer // Tagastab kõik kolmikud stringina.
}

/*
Countdown
	Instructions
Create a function that takes a one-digit integer as input and returns a countdown string. The countdown should start from the given number, skip every second number, and end with 0!. For example, if the input is 7, the function should return "7, 5, 3, 1, 0!".
*/
func Countdown(n int) string { // Defineerib funktsiooni Countdown, mis võtab ühe täisarvu ja tagastab stringi.
	var answer string // Loob tühja stringi muutuja "answer" vastuste hoidmiseks.
	if n == 0 { // Kontrollib, kas sisend on 0.
		return "0!" // Kui sisend on 0, tagastab kohe "0!".
	}
	for i := n; i >= 1; i -= 2 { // Käivitab tsükli alustades sisendist n ja vähendades i väärtust 2 võrra iga sammu järel, kuni i on suurem või võrdne 1-ga.
		answer += string(i+'0') + ", " // Teisendab i väärtuse stringiks, lisades ASCII väärtusele 0, ja lisab selle vastuse stringile koos koma ja tühikuga.
	}
	answer += "0!" // Lisab lõppu "0!".
	return answer // Tagastab lõpliku stringi.
}

/*
Alphanumber
	Instructions
Create a Go function that takes an integer as input and returns a string representing that integer. If the number is negative, preserve the minus sign. Replace the digits with lowercase letters, where 0 becomes a, 1 becomes b, and so on up to 9, which becomes j.
*/

// See funktsioon võtab sisendiks täisarvu ja tagastab selle esindatud stringina, kus numbrid on asendatud tähtedega.
func AlphaNumber(n int) string {
	var answer string // Vastuse jaoks mõeldud string.
	var isNeg bool    // Muutuja, mis näitab, kas arv on negatiivne.

	if n < 0 { // Kui arv on negatiivne,
		isNeg = true // märgib arvu negatiivseks.
		n *= -1      // Teisendab arvu positiivseks töötlemise lihtsustamiseks.
	}

	if n == 0 { // Kui arv on null,
		answer = "a" // määrab vastuseks "a".
	} else { // Kui arv ei ole null,
		for n > 0 { // Kuni arv on suurem kui null,
			digit := n % 10                     // Leiab arvu viimase numbri.
			answer = string('a'+digit) + answer // Teisendab numbri täheks ja lisab selle vastuse algusesse.
			n /= 10                             // Eemaldab arvust viimase numbri.
		}
	}

	if isNeg { // Kui arv oli negatiivne,
		answer = "-" + answer // lisab miinusmärgi vastuse ette.
	}
	return answer // Tagastab vastuse.
}

/*
Simple String to Integer
	Instructions
Create a Go function that converts a valid string representation of a number into an integer.
If the string is not a valid number, the function should return 0. The input will only contain one or more digits, and you don't need to handle signs like + or -.
*/
func SimpleStrToInt(s string) int { // Defineerib funktsiooni SimpleStrToInt, mis võtab ühe stringi ja tagastab täisarvu.
	var answer int // Loob täisarvu tüüpi muutuja "answer", mis hoiab arvutatud tulemust.
	for _, char := range s { // Käib läbi kõik tähemärgid stringis s.
		answer = answer*10 + int(char-'0') // Teisendab tähemärgi arvuks, korrutab praeguse vastuse 10-ga ja lisab uue numbri.
	}
	return answer // Tagastab lõpliku täisarvu.
}

/*
String to Integer
	Instructions
Create a Go function that mimics the behavior of the Atoi function. Atoi converts a string representing a number into an integer.
	The function should return the integer value.
	If the input string is not a valid number, it should return 0.
	The function must handle signs such as + or -.
	The exercise does not require you to return an error.
*/
// See funktsioon teisendab stringi, mis esindab numbrit, täisarvuks.
func StrToInt(s string) int {
	var answer int // Muutuja vastuse jaoks.
	sign := 1      // Muutuja, mis hoiab märki; vaikimisi on see positiivne.
	for i, char := range s {
		if i == 0 { // Käsitseb esimest tähemärki eraldi.
			if char == '-' { // Kui esimene märk on miinus,
				sign = -1 // määra märk negatiivseks.
				continue  // Jätka järgmise tähemärgi juurde.
			} else if char == '+' { // Kui esimene märk on pluss,
				continue // lihtsalt jätkata järgmise tähemärgiga.
			}
		}
		if char < '0' || char > '9' { // Kui tähemärk ei ole vahemikus '0' kuni '9',
			return 0 // tagasta 0, kuna see ei ole kehtiv number.
		}
		answer = answer*10 + int(char-'0') // Arvutab uue väärtuse, korrutades praeguse vastuse 10-ga ja lisades uue numbri väärtuse.
	}
	return answer * sign // Tagastab lõpliku arvu, rakendades märki.
}

func main(){
	fmt.Println(AlphabetMastery(6)) // Prindib tulemuse funktsioonist AlphabetMastery, kus sisendiks on 26 || OUTPUT: abcdef.
	fmt.Println(ReverseAlphabet(5))              // Prindib tähestiku tagurpidi, hüpates üle iga nelja tähe. || OUTPUT: zupkfa
	fmt.Println(Pairs())                         //genereerib stringi kõikidest unikaalsetest paaridest || OUTPUT: "00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99"
	fmt.Println(Combinations())                  // See funktsioon genereerib kõik unikaalsed kolmekohalised kombinatsioonid numbritega 0 kuni 9, kus iga järgnev number on suurem kui eelmine. || OUTPUT: "012, 013, ..., 689, 789"
	fmt.Println(Countdown(7))                    // See funktsioon loob loenduri stringi ühest ühekohalisest numbrist, jättes iga teise numbri vahele, kuni jõuab nullini. OUTPUT: "7, 5, 3, 1, 0!"
	fmt.Println(AlphaNumber(-1280))              //teisendab täisarvu stringiks, kus iga number asendatakse tähestiku vastava tähega, säilitades negatiivse märgi. OUTPUT: "-bcia"
	fmt.Println(SimpleStrToInt("0000000012345")) // Prindib stringi "0000000012345" täisarvuks teisendatud tulemuse.|| OUTPUT:12345
		//StrToInt testcases
	fmt.Println("\nStrToInt testcases")
	fmt.Println(StrToInt("12345"))         // Funktsioon teisendab "12345" täisarvuks 12345. || Output: 12345
	fmt.Println(StrToInt("0000000012345")) // Eesliidetud nullid ei mõjuta numbrilist väärtust, seega "0000000012345" muutub 12345. || Output: 12345
	fmt.Println(StrToInt("012 345"))       // Tühik on kehtetu märk numbristringis, mistõttu tagastatakse 0. || Output: 0
	fmt.Println(StrToInt("Hello World!"))  // Kuna sisend ei ole kehtiv numbriline string, tagastatakse 0. || Output: 0
	fmt.Println(StrToInt("+1234"))         // Plussmärk (+) ees näitab positiivset arvu, seega "+1234" teisendatakse täisarvuks 1234.|| Output: 1234
	fmt.Println(StrToInt("-1234"))         // Miinusmärk (-) ees näitab negatiivset arvu, seega "-1234" teisendatakse täisarvuks -1234. || Output: -1234
	fmt.Println(StrToInt("++1234"))        // Korduvad plussmärgid ei ole kehtivad, mistõttu tagastatakse 0. || Output: 0
	fmt.Println(StrToInt("--1234"))        // Korduvad miinusmärgid ei ole kehtivad, mistõttu tagastatakse 0. || Output: 0

}
