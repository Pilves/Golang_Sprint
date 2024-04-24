package main

import (
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
	var answer []int // Step 1: Initialize an empty slice of integers
	if min >= max {  // Step 2: Handle the case where min is greater than or equal to max
		return nil // Return a nil slice directly
	}
	// Step 3: Populate the slice with integers from min to max-1
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer // Step 4: Return the populated slice
}

/*
Remove Elements In Range
	Instructions
Write a Go function that takes an array of float64, two indices, and removes elements that lie between these indices from the array. The lower index should be removed, and the upper index should be kept. The function should return the resulting modified array. The indices can be negative or larger than the length of the array, but the function should still remove the elements that fall within the specified range. The indices might also be in wrong order.*/

// Funktsioon eemaldab listist teatud vahemikus olevad elemendid. Vahemiku alumine piir eemaldatakse ja ülemine jäetakse.
func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	// Step 1: Correct the order of indices if necessary
	if to < from {
		to, from = from, to
	}
	// Step 2: Adjust indices to be within the valid range
	if from < 0 {
		from = 0 // Set from to 0 if it is negative
	}
	if to > len(arr) {
		to = len(arr) // Set to to the length of the array if it exceeds
	}
	// Step 3: Remove elements between from (inclusive) and to (exclusive)
	var answer []float64
	answer = append(arr[:from], arr[to:]...)
	return answer
}

/*
Balance Out
	Instructions
Write a Go function that takes an array of booleans and adds the minimum number of booleans necessary so that the count of true and false values in the array is equal. The function should return the resulting modified array. The order of the elements must remain the same and new elements should be added at the end of the array.*/

// Funktsioon tasakaalustab tõesete ja väärade väärtuste arvu listis, lisades vajalikud elemendid lõppu.
func BalanceOut(arr []bool) []bool {
	// Step 1: Initialize counters for true and false
	countTrue := 0
	countFalse := 0
	// Step 2: Count the number of true and false in the array
	for _, i := range arr {
		if i == true {
			countTrue++
		} else if i == false {
			countFalse++
		}
	}
	// Step 3: Calculate the difference and identify which element to add
	if countTrue > countFalse {
		// Step 4: Append the required number of booleans to the array
		for i := 0; i < (countTrue - countFalse); i++ {
			arr = append(arr, false)
		}
	} else if countTrue < countFalse {
		// Step 4: Append the required number of booleans to the array
		for i := 0; i < (countFalse - countTrue); i++ {
			arr = append(arr, true)
		}
	}
	// Step 5: Return the modified array
	return arr
}

/*
Filter by Sum
	Instructions
Write a Go function that takes a 2D array of integers and an integer value. The function should filter out all subarrays from the array (2D) in which the sum of elements is lower than the given value. The resulting modified 2D array should be returned.
*/
//Funktsioon filtreerib välja alamlistid, mille elementide summa on väiksem kui etteantud piirväärtus.
func FilterBySum(arr [][]int, limit int) [][]int {
	if len(arr) == 0 {
		return nil
	}
	// Step 1: Initialize an empty list to hold the filtered subarrays
	var result = make([][]int, 0)
	// Step 2: Iterate over each subarray in the matrix
	for _, subarray := range arr {
		// Step 3: Calculate the sum of elements in the current subarray
		sum := 0
		for _, num := range subarray {
			sum += num
		}
		// Step 4: Check if the sum meets the threshold
		if sum >= limit {
			// Step 5: If the sum is greater than or equal to the threshold, include the subarray in the result
			result = append(result, subarray)
		}
	}
	// Step 6: Return the list of filtered subarrays
	return result
}

/*
Sort Integer Table
	Instructions
Write a function that sorts a slice of integers in ascending order.
*/
//Funktsioon sorteerib täisarvude järjendi kasvavas järjekorras.
func SortIntegerTable(table []int) []int {
	// Step 1: Retrieve the number of elements in the slice
	n := len(table)
	// Step 2: Implement Bubble Sort algorithm
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// Step 3: Compare adjacent elements and swap if in wrong order
			if table[j] > table[j+1] {
				// Swap elements at positions j and j+1
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	// Step 4: Return the sorted slice
	return table
}

/*
Bulk Atoi
	Instructions
Write a Go function that takes an array of strings, applies the StrToInt function you wrote earlier to every element in the array, and returns the resulting integer values as a new array. Note that you cannot call the StrToInt function from the package in our current learning environment, you have to copypaste the function instead.
*/
//Funktsioon teisendab stringide listi täisarvudeks, kasutades StrToInt funktsiooni.
func BulkAtoi(arr []string) []int {
	// Step 1: Initialize an empty array for integers
	result := make([]int, 0)
	// Step 2: Loop through each string in the input array
	for _, i := range arr {
		// Step 3: Convert the string to an integer using StrToInt
		num := StrToInt(i)
		// Step 4: Append the integer to the integerArray
		result = append(result, num)
	}
	// Step 5: Return the array of integers
	return result
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
	// Step 1: Handle edge cases
	if n <= 0 {
		return nil
	}
	// Step 2: Initialize containers for results and the current combination
	result := []string{}
	stack := make([]int, n)
	// Step 3: Set initial combination in the stack
	for i := range stack {
		stack[i] = i // Initial combination: 0, 1, 2, ..., n-1
	}
	// Step 4: Iterate to generate all combinations
	for {
		// Step 4.1: Convert the current stack to a string and store it
		comb := ""
		for _, v := range stack {
			comb += fmt.Sprintf("%d", v)
		}
		result = append(result, comb)
		// Step 4.2: Find the rightmost element that can be incremented
		i := n - 1
		for i >= 0 && stack[i] == 9-(n-1-i) {
			i--
		}
		// Step 4.3: Check if the entire combination has been processed
		if i < 0 {
			break // If no element can be incremented, exit the loop
		}
		// Step 4.4: Increment the current element and adjust the following elements
		stack[i]++
		for j := i + 1; j < n; j++ {
			stack[j] = stack[j-1] + 1
		}
	}
	// Step 5: Return all generated combinations
	return result
}

func main() {
	fmt.Println(GenerateRange(5, 10))                                               // Prindib arvud 5 kuni 9.
	fmt.Println(GenerateRange(10, 5))                                               // Prindib tühja järjendi, kuna 10 on suurem kui 5.
	fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1)) // Eemaldab teatud vahemikus olevad elemendid.
	fmt.Println(BalanceOut([]bool{true, false, false, false}))                      // Tasakaalustab tõeste ja väärade väärtuste arvu.
	fmt.Println(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9))           // Filtreerib alamjärjendid, mille summa on alla 9.
	fmt.Println(SortIntegerTable([]int{5, 4, 3, 2, 1, 0}))                          // Sorteerib täisarvud kasvavas järjekorras.
	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))                             // Teisendab stringid täisarvudeks.
	fmt.Println(CombN(1))                                                           // Prindib kõik ühekohalised kombinatsioonid.
	fmt.Println(CombN(8))                                                           // Prindib kõik kolmekohalised kombinatsioonid.
}
