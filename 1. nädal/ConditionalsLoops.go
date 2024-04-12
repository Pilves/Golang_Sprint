package main // See määrab faili paketi, mis on peamine tööpakett.

import "fmt" // Impordib fmt libary, mida kasutatakse andmete väljastamiseks.

/*
Is Negative?
	Instructions
Create a Go function that takes an integer as input and returns a boolean value (true or false) to indicate whether the input integer is negative or not.
	Loo funktsioon, mis võtab sisendiks täisarvu ja tagastab tõeväärtuse (true või false), mis näitab, kas sisend on negatiivne või mitte.
*/
func IsNegative(n int) bool { // Funktsioon kontrollib, kas arv on negatiivne.
	if n < 0 { // Kui arv on väiksem kui 0,
		return true // tagasta true.
	}
	return false // Muidu tagasta false.
}


/*
Int vs Float
	Instructions
Write a Go function that takes an integer and a float as input and returns a string. The function should compare the integer to the float and return one of the following: Integer if the integer is larger, Float if the float is larger, or Same if they are equal.
	Kirjuta funktsioon, mis võtab sisendiks täisarvu ja floati ning tagastab stringi. Funktsioon võrdleb täisarvu ujukomaarvuga ja tagastab ühe järgmistest: "Integer" kui täisarv on suurem, "Float" kui float on suurem või "Same" kui nad on võrdsed.
 */
func IntVsFloat(i int, f float32) string { // Funktsioon võrdleb täisarvu ja ujukomaarvu.
	if float32(i) > f { // Kui täisarv on suurem,
		return "Integer" // tagasta "Integer".
	} else if float32(i) < f { // Kui ujukomaarv on suurem,
		return "Float" // tagasta "Float".
	} else { // Kui väärtused on võrdsed,
		return "Same" // tagasta "Same".
	}
}

/*
Season
	Instructions
While if or if-else statements are very useful, sometimes they are a little uncomfortable. switch statement is very useful if one needs to perform more checks for a value or if there are multiple cases for one check.
In this task you need to make a function that takes a string, that can contain a month name. If a month name is given, the season has to be returned. Otherwise return "invalid input: " with the input appended to it.
	Months:
		winter: jan, feb, dec
		spring: mar, apr, may
		summer: jun, jul, aug
		autumn: sep, oct, nov
*/

// Selgitus: See funktsioon tagastab aastaaja sõltuvalt kuu nimest.
func Season(month string) string { // Defineerib funktsiooni "Season", mis võtab sisendiks kuu nime.

	switch month { // Võrdleb kuu nime ja valib vastava aastaaja.
	case "jan", "feb", "dec": // Kui kuu on jaanuar, veebruar või detsember:
		return "winter" // Tagastab "winter".
	case "mar", "apr", "may": // Kui kuu on mar, apr või may:
		return "spring" // Tagastab "spring".
	case "jun", "jul", "aug": // Kui kuu on jun, jul või aug:
		return "summer" // Tagastab "summer".
	case "sep", "oct", "nov": // Kui kuu on sep, oct või nov:
		return "autumn" // Tagastab "autumn".
	default: // Kui kuu nimi ei vasta ühelegi eelnevale:
		return "invalid input: " + month // Tagastab veateate koos sisendiga.
	}
}

/*
Accumulate
	Instructions
Create a Go function that takes a non-negative integer n. If n is positive, return the sum of all the digits from 0 to n, including n itself. If n is negative, return 0.
For example, in case of 4, the sum would be 0 + 1 + 2 + 3 + 4 = 10.
*/

// See funktsioon arvutab numbrite summa nullist kuni n-ni, kaasa arvatud n. Kui n on negatiivne, tagastab 0.
func Accumulate(n int) int {
	answer := 0               // Algusväärtus vastusele.
	for i := n; i >= 0; i-- { // Käib läbi numbrid n-st nullini.
		answer += i // Lisab iga numbri vastusele.
	}
	return answer // Tagastab summa.
}

/*
Between Limits
	Instructions
Build a function that accepts two runes as input and returns a string containing all the runes that come between these two runes in the alphabet, in the correct order. For instance, if the input runes are 'f' and 'j', the function should return "ghi" regardless of the order in which the runes are given.
	Ehita funktsioon, mis võtab sisendiks kaks rune-tüüpi tähte ja tagastab stringi, mis sisaldab kõiki tähestiku tähti nende kahe tähe vahel õiges järjestuses. Näiteks, kui sisendtähed on 'f' ja 'j', peaks funktsioon tagastama "ghi" olenemata tähtede sisestamise järjekorrast.
*/
func BetweenLimits(from, to rune) string { // Funktsioon leiab kõik tähed kahe rune vahel.
	var answer string
	if from > to { // Kui esimene täht on suurem kui teine,
		to, from = from, to // vahetab tähti.
	}
	for i := from + 1; i < to; i++ { // Käib läbi kõik tähed esimese ja teise tähe vahel.
		answer += string(i) // Lisab iga tähe vastusele.
	}
	return answer // Tagastab lõpliku stringi.
}

/*
Do Operation
	Instructions
You're tasked with creating a function called Doop. This function takes three parameters:
	A value (integer).
	An operator, which can be one of the following: +, -, /, *, %.
	Another value (integer).
	In case of an invalid operator, invalid values, or incorrect number of arguments, the function returns 0. The program must also handle modulo and division by 0.
*/
func Doop(a int, op string, b int) int { // Funktsioon teostab aritmeetilisi operatsioone kahe arvu vahel.
	switch op { // Kontrollib operaatorit.
	case "+": // Liitmine.
		return a + b
	case "-": // Lahutamine.
		return a - b
	case "/": // Jagamine.
		if b == 0 { // Kui jagaja on null,
			return 0 // tagasta 0.
		}
		return a / b
	case "*": // Korrutamine.
		return a * b
	case "%": // Modulo.
		if b == 0 { // Kui modulo jagaja on null,
			return 0 // tagasta 0.
		}
		return a % b
	default: // Kui operaator ei ole kehtiv,
		return 0 // tagasta 0.
	}
}

/*
Is Leap Year?
	Instructions
Write a Go function that takes an integer representing a year and returns a boolean value indicating whether that year is a leap year or not. A leap year is a year that is evenly divisible by 4, except for years that are divisible by 100 but not by 400.
	Liigaasta on aasta, mis jagub ühtlaselt 4-ga, välja arvatud aastad, mis jaguvad 100-ga, kuid mitte 400-ga.
*/
func IsLeapYear(year int) bool { // Kontrollib, kas aasta on liigaasta.
	if year%4 == 0 && year%100 != 0 || year%400 == 0 { // Liigaasta reeglid.
		return true // Tagasta true, kui aasta on liigaasta.
	}
	return false // Muidu tagasta false.
}

/*
Count Divisible
	Instructions
Write a Go function that takes four integers as input: from, to, step, and divisor. The function should check for exceptions and return 0 if any of the following conditions are met:
	step is negative or zero.
	divisor is zero.
	Otherwise, it should loop through the range of numbers from from (inclusive) to to (exclusive), checking every step-th element, and return the count of elements among these that are divisible by the divisor.
	funktsioon peab läbima loop-i numbrite vahemikus from(kaasa arvatud) kuni to(välja arvatud) ja tagastama mitu numbrit seal vahemikus jaguvad divisoriga
	*/
func CountDivisible(from, to, step, divisor int) int { // Loendab, mitu arvu jagub kindla arvuga antud vahemikus.
	var count int
	if step <= 0 || divisor == 0 { // Kontrollib, kas samm või jagaja on kehtetu.
		return 0 // Kui jah, tagasta 0.
	}
	for i := from; i < to; i += step { // Käib läbi numbrid algusest lõpuni kindla sammuga.
		if i%divisor == 0 { // Kui number jagub jagajaga,
			count++ // suurenda loendurit.
		}
	}
	return count // Tagasta loenduri väärtus.
}

/*
Find Dividend
	Instructions
Write a Go function that takes three integers as input: from, to, and divisor. The function should loop through the numbers from from (inclusive) to to (exclusive) and return the first number in that range that is divisible by the divisor. If there is no such number, the function should return -1.
	Kirjuta funktsioon, mis võtab sisendiks kolm täisarvu: from, to ja divisor. Funktsioon peab läbima loop-i numbrite vahemikus from(kaasa arvatud) kuni to(välja arvatud) ja tagastama esimese numbri selles vahemikus, mis jagub jagajaga. Kui sellist numbrit ei ole, peab funktsioon tagastama -1.
*/
func FindDividend(from, to, divisor int) int { // Leiab esimese jagatava arvu vahemikus.
	for i := from; i < to; i++ { // Käib läbi numbrid algusest lõpuni.
		if i%divisor == 0 { // Kui number jagub jagajaga,
			return i // tagasta see number.
		}
	}
	return -1 // Kui ükski number ei jagu, tagasta -1.
}

func main() { // Peamine funktsioon, mis testib ülaltoodud funktsioone.
	fmt.Println(IsNegative(6)) // Kontrollib, kas 6 on negatiivne || OUTPUT: false.
	fmt.Println(IntVsFloat(5, 8.8)) // Võrdleb 5 ja 8.8 || OUTPUT: Float.
	fmt.Println(Season("feb")) // Tagastab veebruari aastaaja || OUTPUT: winter.
	fmt.Println(Season("September")) // Tagastab "September" aastaaja || OUTPUT: invalid input: September.
	fmt.Println(Accumulate(4)) // Arvutab summa nullist kuni 4 || OUTPUT: 10.
	fmt.Println(BetweenLimits('j', 'f')) // Leiab tähed vahemikus 'j' kuni 'f' || OUTPUT: ghi.
	fmt.Println(Doop(5, "+", 3)) // Teostab operatsiooni 5 + 3 || OUTPUT: 8.
	fmt.Println(Doop(8, "/", 0)) // Teostab operatsiooni 8 jagatud 0-ga || OUTPUT: 0.
	fmt.Println(IsLeapYear(2020)) // Kontrollib, kas 2020 on liigaasta || OUTPUT: true.
	fmt.Println(CountDivisible(5, 17, 2, 3)) // Loendab jaguvad numbrid vahemikus 5 kuni 17 || OUTPUT: 2.
	fmt.Println(FindDividend(5, 17, 4)) // Leiab esimese jagatava numbri vahemikus 5 kuni 17 || OUTPUT: 8.
	fmt.Println(FindDividend(10, 18, 9)) // Leiab esimese jagatava numbri vahemikus 10 kuni 18 || OUTPUT: -1.
}
