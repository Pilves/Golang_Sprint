package main // Määrab, et selle koodi package.

import ( // Importib vajalikud libaryd.
	"fmt"  // Võimaldab kasutada väljundfunktsioone.
	"math" // Võimaldab kasutada matemaatilisi funktsioone.
)

/*
Abacus
	Instructions
Write a function that takes two integer (int) values as input and returns the result of integer division of the first value by the second.
	Kirjuta funktsioon, mis võtab kaks täisarvu (int) ja tagastab nende jagamise tulemuse täisarvuna.
*/

func Abacus(a int, b int) int { // Funktsioon võtab kaks täisarvu ja jagab esimese teisega.
	return a / b // Tagastab jagamise tulemuse.
}

/*
Mean
	Instructions
Write a function that takes three float32 values as input and returns their mean as a float32 value.
	Kirjuta funktsioon, mis võtab kolm float32 tüüpi arvu ja tagastab nende keskmise float32 tüüpi arvuna.
*/

func Mean(a, b, c float32) float32 { // Funktsioon võtab kolm arvu ja arvutab keskmise.
	return (a + b + c) / 3 // Tagastab kolme arvu keskmise.
}

/*
Casting
	Instructions
Write a function that accepts a float64 value, rounds it to the nearest integer, and then returns the result as an int.
	For this task you're allowed to use:
		math
Kirjuta funktsioon, mis võtab ühe float64 tüüpi arvu, ümardab selle lähima täisarvuni ja tagastab tulemuse int tüüpi arvuna.*/

func Casting(n float64) int { // Funktsioon ümardab arvu lähima täisarvuni.
	return int(math.Round(n)) // Kasutab math.Round funktsiooni ja teisendab tulemuse täisarvuks.
}

/*
Shift By
	Instructions
Create a Go function that takes a single lowercase letter as a rune and an int 'step.' The function should shift the letter in the alphabet by the specified 'step' value, and return the resulting letter. For example, if you shift 'a' by 4 steps, it should become 'e.' Remember to handle looping back to the start of the alphabet if needed.
Only lower case characters are tested.

	Loo funktsioon, mis võtab ühe väiketähe (rune tüüpi) ja sammu (int tüüpi). Funktsioon peaks nihutama tähte tähestikus määratud sammude võrra ja tagastama tulemuse.
	Näiteks, kui nihutad 'a' 4 sammu, peaks tulema 'e'.
*/

func ShiftBy(r rune, step int) rune { // Funktsioon nihutab tähte tähestikus etteantud sammude võrra.
	return rune((int(r-'a')+step)%26 + 'a') // Arvutab uue tähe positsiooni ja tagastab selle.
}

/*
String Concatenation

	Instructions

Create a function that takes two strings and a delimiter. The function should combine the two strings, placing the delimiter between them, and return the combined result as a single string.

	Loo funktsioon, mis võtab kaks stringi ja eraldaja. Funktsioon peaks ühendama mõlemad stringid, paigutades eraldaja nende vahele, ja tagastama ühendatud tulemuse ühe stringina.
*/
func StrConcat(s1, s2, delim string) string { // Funktsioon ühendab kaks stringi, lisades nende vahele eraldaja.
	return s1 + delim + s2 // Tagastab ühendatud stringi.
}

/*
Reverse Alphabet Value

	Instructions
Write a Go function that takes a lowercase letter rune ('a'-'z') as input and returns its reverse letter in the Latin alphabet as rune. For example, 'a' should be transformed to 'z', 'y' should become 'b', and so on. The function should maintain the case of the input letter.
	Kirjuta funktsioon, mis võtab väiketähe rune ('a'-'z') ja tagastab selle tähestiku vastandtähe rune'ina.
*/

func ReverseAlphabetValue(ch rune) rune { // Funktsioon leiab tähe vastandtähe tähestikus.
	return 'a' + 'z' - ch // Arvutab vastandtähe ja tagastab selle.
}

func main() { // Peamine funktsioon, kus testitakse ülaldefineeritud funktsioone.
	fmt.Println(Abacus(8, 3))                       // Prindib 8 jagatud 3 || OUTPUT: 2.
	fmt.Println(Mean(1.15, -2.0, 8.35))             // Arvutab keskmise kolmest arvust || OUTPUT: umbes 2.50.
	fmt.Println(Casting(1.15))                      // Ümardab 1.15 lähima täisarvuni || OUTPUT: 1.
	fmt.Println(string(ShiftBy('a', 4)))            // Nihutab 'a' neli sammu || OUTPUT: 'e'.
	fmt.Println(StrConcat("Hello", "World!", ", ")) // Ühendab "Hello" ja "World!" komaga || OUTPUT: "Hello, World!".
	fmt.Println(string(ReverseAlphabetValue('a')))  // Leiab 'a' vastandtähe || OUTPUT: 'z'.
}
