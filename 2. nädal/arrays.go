package main

import(
	"fmt"
)

/*
Generate Range
	Instructions
Create a Go function that takes two integers, min and max, as input. The function should return a slice of integers containing all values between min (inclusive) and max (exclusive) without using the make function to create the slice.
If min is greater than or equal to max, the function should return a nil slice.
Funktsioon võtab sisendiks kaks täisarvu, min ja max, ning tagastab nende kahe arvu vahelise arvude jada.
Kui min on suurem või võrdne max'iga, tagastatakse tühi jada.
*/

func GenerateRange(min, max int) []int {
	var answer []int // Vastuseks on arvude jada.
	if min >= max{
		return []int{} // Kui min on suurem või võrdne max'iga, tagastame tühja jada.
	}
	for i:= min; i <max;i++{
		answer = append(answer, i) // Lisame iga arvu järjest vastusesse.
	}
	return answer // Tagastame arvude jada.
}


/*
Remove Elements In Range
	Instructions
Write a Go function that takes an array of float64, two indices, and removes elements that lie between these indices from the array. The lower index should be removed, and the upper index should be kept. The function should return the resulting modified array. The indices can be negative or larger than the length of the array, but the function should still remove the elements that fall within the specified range. The indices might also be in wrong order.*/

//Funktsioon eemaldab listist teatud vahemikus olevad elemendid. Vahemiku alumine piir eemaldatakse ja ülemine jäetakse.
func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	length := len(arr) // Määrame listi pikkuse.
	if from < 0{
		from += length
	}
	if to < 0{
		to += length
	}
	if from < 0{
		from = 0
	}
	if to > length{
		to = length
	}
	if from > to{
		from, to = to, from // Kui indeksid on vale järjekorras, vahetame neid.
	}
	var result []float64
	if from >= 0 && to <= length && from < to{
		result = append(arr[:from], arr[to:]...) // Ühendame osad, mis ei jää vahemikku.
	}else{
		result = arr
	}
	return result // Tagastame muudetud listi.

}
/*
Balance Out
	Instructions
Write a Go function that takes an array of booleans and adds the minimum number of booleans necessary so that the count of true and false values in the array is equal. The function should return the resulting modified array. The order of the elements must remain the same and new elements should be added at the end of the array.*/

//Funktsioon tasakaalustab tõesete ja väärade väärtuste arvu listis, lisades vajalikud elemendid lõppu.
func BalanceOut(arr []bool) []bool {
	trueCount :=0 // Loendame tõeste väärtuste arvu.
	falseCount := 0 // Loendame väärade väärtuste arvu.
	for _,value := range arr{
		if value{
			trueCount++
		}else{
			falseCount++
		}
	}

	if trueCount > falseCount{
		for i:=0; i< (trueCount-falseCount);i++{
			arr = append(arr, false) // Lisame vajaliku arvu väärade väärtuste.
		}
	}else if falseCount > trueCount{
		for i :=0; i< (falseCount-trueCount);i++{
			arr = append(arr, true) // Lisame vajaliku arvu tõeseid väärtusi.
		}
	}
	return arr // Tagastame muudetud listi.
}

/*
Filter by Sum
	Instructions
Write a Go function that takes a 2D array of integers and an integer value. The function should filter out all subarrays from the array (2D) in which the sum of elements is lower than the given value. The resulting modified 2D array should be returned.
*/
//Funktsioon filtreerib välja alamlistid, mille elementide summa on väiksem kui etteantud piirväärtus.
func FilterBySum(arr [][]int, limit int) [][]int {
	var result [][]int
	for _,subarray := range arr{
		sum :=0
		for _, num := range subarray{
			sum += num // Liidame kõik numbrid alamlistis.
		}
		if sum >=limit{
			result = append(result, subarray) // Kui summa on piisav, lisame tulemusele.
		}
	}
	return result // Tagastame filtreeritud listi.
}


/*
Sort Integer Table
	Instructions
Write a function that sorts a slice of integers in ascending order.
*/
//Funktsioon sorteerib täisarvude järjendi kasvavas järjekorras.
func SortIntegerTable(table []int) []int {
	n:= len(table) // Täisarvude listi pikkus.
	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			if table[j] > table[j+1]{
				table[j], table[j+1] = table[j+1], table[j] // Vahetame elemendid, kui need on vales järjekorras.
			}
		}
	}
	return table // Tagastame sorteeritud listi.
}


/*
Bulk Atoi
	Instructions
Write a Go function that takes an array of strings, applies the StrToInt function you wrote earlier to every element in the array, and returns the resulting integer values as a new array. Note that you cannot call the StrToInt function from the package in our current learning environment, you have to copypaste the function instead.
*/
//Funktsioon teisendab stringide listi täisarvudeks, kasutades StrToInt funktsiooni.
func BulkAtoi(arr []string) []int {
	result := make([]int, len(arr)) // Loome lõpplisti tulemustele.
	for i, str := range arr{
		result[i] = StrToInt(str) // Teisendame iga stringi täisarvuks.
	}
	return result // Tagastame täisarvude listi.
}

func StrToInt(s string) int {
	var result int
	var sign = 1

	for i, char := range s {
		if i == 0 {
			if char == '-' {
				sign = -1
				continue
			} else if char == '+' {
				continue
			}

		}

		if char < '0' || char > '9' {
			return 0
		} else {
			result = result*10 + int(char-'0')
		}
	}

	return result * sign

}

/*
Combination of N
	Instructions
Dive into the world of number combinations! Your goal is to create a function that reveals all possible combinations of ascending digits of length n.
*/
//Funktsioon loob kõik võimalikud kombinatsioonid n pikkusega, kasvavas järjekorras.
func CombN(n int) []string {
	if n <= 0{
		return []string{} // Kui n on 0 või negatiivne, tagastame tühja listi.
	}
	results:= []string{}
	stack := make([]int,n)
	var startIndex, num int

	for{
		comb :=""
		for i:=0;i<n;i++{
			comb += fmt.Sprintf("%d", stack[i]) // Loome kombinatsiooni stringi.
		}
		results = append(results, comb)
		for startIndex = n-1; startIndex >= 0; startIndex--{
			if stack[startIndex] < 9-(n-1-startIndex){
				num = stack[startIndex] +1
				break
			}
		}
		if startIndex < 0{
			break
		}
		for i := startIndex;i<n;i++{
			stack[i]=num
			num++
		}
	}
	return results // Tagastame kõik kombinatsioonid.
}

func main(){
	fmt.Println(GenerateRange(5, 10)) // Prindib arvud 5 kuni 9.
	fmt.Println(GenerateRange(10, 5)) // Prindib tühja järjendi, kuna 10 on suurem kui 5.
	fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1)) // Eemaldab teatud vahemikus olevad elemendid.
	fmt.Println(BalanceOut([]bool{true, false, false, false})) // Tasakaalustab tõeste ja väärade väärtuste arvu.
	fmt.Println(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9)) // Filtreerib alamjärjendid, mille summa on alla 9.
	fmt.Println(SortIntegerTable([]int{5, 4, 3, 2, 1, 0})) // Sorteerib täisarvud kasvavas järjekorras.
	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"})) // Teisendab stringid täisarvudeks.
	fmt.Println(CombN(1)) // Prindib kõik ühekohalised kombinatsioonid.
	fmt.Println(CombN(3)) // Prindib kõik kolmekohalised kombinatsioonid.
}
