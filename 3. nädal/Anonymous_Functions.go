package main

import "fmt"

/*
Array Any
	Instructions
You're tasked with creating a function ArrAny that returns true for a string slice if, when that string slice is passed through an f function, at least one element returns true.

For this task you will need to have in the same file some functions you have previously created:
	IsLower
	IsNumeric
You will also need to implement two new functions:
	IsUpper Returns true if the string given contains only uppercase characters, and false otherwise.
	IsAlphanumeric Returns true if the string contains only alphanumeric characters, and false otherwise.
*/

func ArrAny(f func(string) bool, a []string) bool {
	// Iterate through each element in the array
	for _, str := range a{
		// Call the provided function 'f' with the current element 'str'
		if f(str){
			// If the function returns true for any element, return true
			return true
		}
	}
	// If no element in the array satisfies the condition, return false
	return false
}


func IsUpper(s string) bool {
	// Iterate through each character in the string
	for _, char := range s{
		// If the character is not in uppercase range (A-Z), return false
		if char < 'A' || char > 'Z'{
			return false
		}
	}
	// If all characters are uppercase, return true
	return true
}


func IsAlphanumeric(s string) bool {
	// Iterate through each character in the string
	for _, char := range s {
		// If the character is not alphanumeric, return false
		if (char < '0' || char > '9') && (char < 'A' || char > 'Z') && (char < 'a' || char > 'z'){
			return false
		}
	}
	// If all characters are alphanumeric, return true
	return true
}


func IsNumeric(s string) bool {
	for _, char:= range s{
		if char < '0' || char > '9'{
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _,char := range s{
		if char < 'a' || char > 'z'{
			return false
		}
	}
	return true
}

/*
Array Count If
	Instructions
You're tasked with creating a function ArrCountIf that returns the number of elements in a string slice for which the f function returns true.

For this task you will need to have in the same file some functions you have previously created:
	IsLower
	IsUpper
	IsNumeric
	IsAlphanumeric
*/

func ArrCountIf(f func(string) bool, tab []string) int {
	// Initialize a counter to keep track of matching elements
	count := 0
	// Iterate through each element in the array
	for _, str := range tab{
		// Call the provided function 'f' with the current element 'str'
		if f(str){
			// If the function returns true for the current element, increment the counter
			count++
		}
	}
	// Return the total count of elements that satisfy the condition
	return count
}

/*
Array Map
	Instructions
You're tasked with creating a function ArrMap that, for an int slice, applies a function of this type func(int) bool on each element of that slice and returns a slice of all the return values.

For this task you will need to have in the same file some functions you have previously created:
	IsNegative
	IsPrime
*/

func ArrMap(f func(int) bool, a []int) []bool {
	// Initialize an empty array to store the mapped values
	answer := []bool{}
	// Iterate through each element in the array
	for _, num := range a{
		// Call the provided function 'f' with the current element 'num' and append the result to the answer array
		answer = append(answer, f(num))
	}
	// Return the array containing mapped values
	return answer
}

func IsNegative(n int) bool {
	if n < 0 {
		return true
	}
	return false
}

func IsPrime(n int) bool{
	if n <= 1{
		return false
	}
	for i := 2; i*i <= n; i++{
		if n % i == 0{
			return false
		}
	}
	return true
}

/*
Advanced Sort Word Array
	Instructions
You're tasked with creating a function called AdvancedSortWordArr that sorts a slice of string based on the provided comparison function f.

For this task you will need to have in the same file some functions you have previously created:
	StrCompare
*/

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	// Iterate through each element in the array
	for i:= 0; i< len(a)-1; i++{
		// Iterate through each element except the last i elements
		for j:= 0; j< len(a)-1-i;j++{
			// Compare adjacent elements using the provided comparison function 'f'
			if f(a[j], a[j+1]) > 0{
				// If the comparison result is greater than 0, swap the elements
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	// Return the sorted array
	return a
}

func StrCompare(a, b string) int {
	for i:= 0; i < len(a) && i< len(b);i++{
		if a[i] < b[i]{
			return -1
		}else if a[i]> b[i]{
			return 1
		}
	}
	switch{
		case len(a) < len(b):
			return -1
		case len(a) > len(b):
			return 1
	}
	return 0
}

/*
Is Sorted?
	Instructions
You're tasked with creating a function called IsSorted that returns true if the slice of string is sorted according to a provided comparison function, and false otherwise.
The comparison function is passed as an argument func(a, b string) int, and it should return a positive int if the first argument is greater than the second argument, 0 if they are equal, and a negative int otherwise.
Note that the slice can be sorted in both ascending and descending order.

For this task you will need to have in the same file some functions you have previously created:
	StrCompare
*/

func IsSorted(f func(a, b string) int, arr []string) bool {
	// Initialize variables to track ascending and descending orders
	asc, desc := false, false
	// Iterate through the array starting from the second element
	for i:= 1; i< len(arr); i++ {
		// Compare the current element with the previous one using the provided comparison function 'f'
		comp := f(arr[i-1], arr[i])
		// Check if the comparison indicates descending order
		if comp > 0{
			// If the comparison result is greater than 0, set 'desc' flag to true and break the loop
			desc = true
			break
			// Check if the comparison indicates ascending order
		}else if comp < 0{
			// If the comparison result is less than 0, set 'asc' flag to true and break the loop
			asc = true
			break
		}
	}
	// Iterate through the array starting from the second element again
	for i := 1; i < len(arr); i++{
		// Compare the current element with the previous one using the provided comparison function 'f'
		comp := f(arr[i-1], arr[i])
		// Check if the current ordering violates the detected order (ascending or descending)
		if ( asc && comp > 0) || (desc && comp < 0){
			// If the current ordering violates the detected order, return false
			return false
		}
	}
	// If the array is sorted according to the detected order, return true
	return true
}

func main(){
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}
	fmt.Println(ArrAny(IsNumeric, a1)) // false
	fmt.Println(ArrAny(IsNumeric, a2)) // true
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	fmt.Println(ArrCountIf(IsNumeric, tab1)) // 0
	fmt.Println(ArrCountIf(IsNumeric, tab2)) // 2
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ArrMap(IsPrime, a)) // [false true true false true false]
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	fmt.Println(AdvancedSortWordArr(result, StrCompare)) // [1 2 3 A B C a b c]
	b1 := []string{"a", "b", "c"}
	b2 := []string{"c", "a", "b"}
	fmt.Println(IsSorted(StrCompare, b1)) // true
	fmt.Println(IsSorted(StrCompare, b2)) // false



}
