package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
GCD
	Instructions
Write a Go function that takes two integers as input and returns their greatest common divisor (GCD). The GCD is the largest positive integer that divides both of the input integers without leaving a remainder.
*/

func GCD(a, b int) int {
	// Ensure a is greater than or equal to b
	if b > a {
		// Swap a and b if b is greater
		a, b = b, a
	}
	// Iterate until b becomes 0
	for b != 0 {
		// Update a to b and b to the remainder of a divided by b
		a, b = b, a%b
	}
	// Return the greatest common divisor, which is stored in a
	return a
}

/*
LCM
	Instructions
Write a Go function that takes two integers as input and returns their least common multiple (LCM). The LCM is the smallest positive integer that is divisible by both of the input integers.
*/

func LCM(a, b int) int {
	// Calculate the least common multiple (LCM) using the formula (a * b) / GCD(a, b)
	lcm := (a * b) / GCD(a, b)
	// Return the calculated LCM
	return lcm
}

/*
Is Palindrome?
	Instructions
Write a Go function that takes a string as input and returns a boolean value indicating whether the input string is a palindrome or not. A palindrome is a string that reads the same forwards and backward, considering character case and white spaces.
*/

func IsPalindrome(s string) bool {
	// Initialize variables i and j for indexing from the start and end of the string, respectively
	j := len(s) - 1
	i := 0
	// Iterate until i reaches or surpasses j
	for i < j {
		// Check if characters at positions i and j are not equal
		if s[i] != s[j] {
			// If not equal, return false indicating that the string is not a palindrome
			return false
		}
		// Move to the next character from the start and end of the string
		i++
		j--
	}
	// If the loop completes without finding any non-matching characters, return true indicating a palindrome
	return true
}

/*
Are Anagrams?
	Instructions
Write a Go function that takes two strings as input and returns a boolean value indicating whether the input strings are anagrams or not. Anagrams are words or phrases formed by rearranging the letters of another, and in this task, you should ignore differences in character case and whitespace.
For this task you're allowed to use:
	sort
	strings
*/

func AreAnagrams(str1, str2 string) bool {
	// Convert both strings to lowercase and remove spaces
	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))
	// Convert each string to a list of characters
	str1Char := strings.Split(str1, "")
	str2Char := strings.Split(str2, "")
	// Sort the lists of characters
	sort.Strings(str1Char)
	sort.Strings(str2Char)
	// Join the sorted characters back into strings
	str1Sorted := strings.Join(str1Char, "")
	str2Sorted := strings.Join(str2Char, "")
	// Check if the sorted strings are equal, indicating that they are anagrams
	return str1Sorted == str2Sorted
}

/*
Sort Word Array
	Instructions
You're tasked with creating a function called SortWordArr that sorts a string slice in ascending order based on ASCII values.
*/

func SortWordArr(a []string) []string {
	// Iterate through the array from the beginning to the second last element
	for i := 0; i < len(a); i++ {
		// Iterate through the array from the beginning to the element before the last i elements
		for j := 0; j < len(a)-i-1; j++ {
			// Compare adjacent elements and swap if necessary to sort in ascending order
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	// Return the sorted array
	return a
}

/*
String Compressor
	Instructions
Write a Go function that takes a string as input and compresses it by replacing consecutive repeating characters with a count of repetitions followed by the character. However, if there is only one character in a sequence, it should not be prepended with a number. For example, "aaabbaac" should be compressed to "3a2b2ac", but "abcde" should remain unchanged as "abcde".
For this task you're allowed to use:
	fmt
*/

func StrCompress(input string) string {
	n := len(input)
	// If the length of the input string is 0 or 1, return the input string as it is
	if n <= 1 {
		return input
	}
	// Initialize an empty string to store the compressed string
	compressed := ""
	// Initialize a count variable to keep track of the consecutive characters
	count := 1
	// Iterate through the input string starting from the second character
	for i := 1; i < n; i++ {
		// If the current character is equal to the previous character, increment the count
		if input[i] == input[i-1] {
			count++
		} else {
			// If the current character is different from the previous one:
			// If the count is greater than 1, append the count to the compressed string
			if count > 1 {
				compressed += fmt.Sprintf("%d", count)
			}
			// Append the previous character to the compressed string
			compressed += string(input[i-1])
			// Reset the count to 1 for the new character
			count = 1
		}
	}
	// Handle the last character and its count
	// If the count is greater than 1, append the count to the compressed string
	if count > 1 {
		compressed += fmt.Sprintf("%d", count)
	}
	// Append the last character to the compressed string
	compressed += string(input[n-1])
	// Return the compressed string
	return compressed
}

/*
Remove Duplicates
	Instructions
Write a Go function that takes an array of integers as input and returns the same array with duplicate elements removed, preserving the order of the remaining elements. For example, if the input array is [3, 5, 2, 3, 8, 5, 2], the function should return [3, 5, 2, 8].
*/

func RemoveDuplicates(arr []int) []int {
	// Initialize a map to keep track of unique elements
	exists := map[int]bool{}
	// Initialize an empty array to store the result without duplicates
	result := []int{}
	// Iterate through the input array
	for _, num := range arr {
		// Check if the current number is not already present in the map
		if !exists[num] {
			// Append the current number to the result array
			result = append(result, num)
			// Mark the current number as existing in the map
			exists[num] = true
		}
	}
	// Return the result array without duplicates
	return result
}

/*
Longest Climb
	Instructions
Write a Go function that takes an array of integers and returns the longest contiguous subarray in the array in which the elements are increasing. The function should return the contiguous subarray.
*/

func LongestClimb(arr []int) []int {
	n := len(arr)
	// If the array is empty, return nil
	if n == 0 {
		return nil
	}
	// Initialize variables to track the longest climb
	start := 0
	maxLen := 1
	maxStart := 0

	// Iterate through the array starting from the second element
	for i := 1; i < n; i++ {
		// Check if the current element is greater than the previous element
		if arr[i] > arr[i-1] {
			// Calculate the length of the current climb
			currentLen := i - start + 1
			// Update maxLen and maxStart if the current climb is longer
			if currentLen > maxLen {
				maxLen = currentLen
				maxStart = start
			}
		} else {
			// If the current element is not greater than the previous one, update the start index
			start = i
		}
	}
	// Return the longest climb, which is a subarray of the original array
	return arr[maxStart : maxStart+maxLen]
}

/*
Payout
	Instructions
Write a Go function that takes an integer amount and a slice of integers denominations. The function should use the values in denominations to pay exactly the amount. The denominations can be used any number of times, but higher denominations should be preferred. The function should return the resulting denominations as an array of integers in descending order. If the payout cannot be made, return an empty array.
*/

func Payout(amount int, denominations []int) []int {
	// Initialize an empty array to store the payout denominations
	payout := make([]int, 0)
	// Sort the denominations array in descending order
	n := len(denominations)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if denominations[j] < denominations[j+1] {
				denominations[j], denominations[j+1] = denominations[j+1], denominations[j]
			}
		}
	}
	// Iterate through the sorted denominations
	for _, denom := range denominations {
		// Check if the current denomination can be used for the payout
		for amount >= denom {
			// Add the current denomination to the payout array
			payout = append(payout, denom)
			// Update the remaining amount
			amount -= denom
		}
	}
	// If the entire amount has been paid out, return the payout array
	if amount == 0 {
		return payout
	}
	// If there is still remaining amount, return an empty array
	return []int{}
}

/*
From Roman
	Instructions
Write a Go function that takes a valid Roman numeral as input and converts it into an integer. The function should return the integer representation of the valid Roman numeral input.
*/

func FromRoman(s string) int {
	// Define the values of Roman numerals
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// Initialize the result to 0
	result := 0

	// Iterate through the characters of the input string
	for i := 0; i < len(s); i++ {
		// Get the value of the current Roman numeral
		value := romanValues[s[i]]
		// Check if there is a next character and its value is greater than the current one
		if i+1 < len(s) && romanValues[s[i+1]] > value {
			// Subtract the current value from the result
			result -= value
		} else {
			// Add the current value to the result
			result += value
		}
	}
	// Return the resulting integer
	return result
}

/*
To Roman
	Instructions
Write a Go function that takes an integer and converts it into a Roman numeral. The function should return "Invalid input" if the input integer is less than 1 or more than 3999. Otherwise, it should return the Roman numeral representation of the valid input integer.
*/

func ToRoman(num int) string {
	// Define the mapping of integers to Roman numerals
	romanNumerals := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	// Define the values of Roman numerals in descending order
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	// Check if the input number is out of the valid range
	if num < 1 || num > 3900 {
		return "Invalid input"
	}
	// Initialize an empty string to store the Roman numeral representation
	roman := ""

	// Iterate through the values of Roman numerals
	for _, value := range values {
		// Repeat adding the corresponding Roman numeral while num is greater than or equal to the current value
		for num >= value {
			// Append the Roman numeral to the result string
			roman += romanNumerals[value]
			// Update num by subtracting the current value
			num -= value
		}
	}
	// Return the Roman numeral representation of the input number
	return roman
}

/*
Pascal's Triangle
	Instructions
Write a Go function that takes an integer as input and returns a 2D integer array representing Pascal's triangle up to the specified number of levels. Pascal's triangle is a triangular array of binomial coefficients where each number is the sum of the two directly above it. The function should generate the specified number of levels of Pascal's triangle.
*/

func PascalsTriangle(n int) [][]int {
	// Initialize an empty list to represent Pascal's triangle
	triangle := make([][]int, n)

	// Iterate through rows of the triangle
	for i := range triangle {
		// Create a new row with i + 1 elements
		triangle[i] = make([]int, i+1)
		// Set the first and last elements of each row to 1
		triangle[i][0] = 1
		triangle[i][i] = 1
	}
	// Populate the inner elements of the triangle
	for i := 2; i < n; i++ {
		for j := 1; j < len(triangle[i])-1; j++ {
			// Calculate the value of the current element based on the values of the previous row
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	// Return the completed Pascal's triangle
	return triangle
}

func main() {
	fmt.Println(GCD(12, 18))                          // 6
	fmt.Println(LCM(12, 18))                          // 36
	fmt.Println(IsPalindrome("level"))                //true
	fmt.Println(AreAnagrams("Listen", "S i l e n t")) // true
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	fmt.Println(SortWordArr(result))                                      // [1 2 3 A B C a b c]
	fmt.Println(StrCompress("aaabbaac"))                                  // "3a2b2ac"
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 2, 4, 8, 8, 1, 2, 0, 8})) // [1, 2, 3, 4, 8, 0]
	fmt.Println(LongestClimb([]int{8, 4, 2, 1, 2, 4, 8, 2, 4, 8}))        // [1 2 4 8]
	amount := 128
	denominations := []int{1, 2, 5, 10, 20, 50, 100, 200}
	fmt.Println(Payout(amount, denominations)) // [100 20 5 2 1]
	fmt.Println(FromRoman("CXXVIII"))          // 128
	fmt.Println(ToRoman(128))                  // "CXXVIII"
	fmt.Println(PascalsTriangle(2))            // [[1] [1 1]]

}
