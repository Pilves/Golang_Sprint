package main

import(
	"fmt"
	"strings"
	"math"
)

/*
Get first rune
	Instructions
Create a Go function that takes a non-empty string as input and returns the first character of the string as a rune.
*/
func GetFirstRune(s string) rune {
	// Initialize a variable 'answer' to store the rune result
	var answer rune
	// Initialize a variable 'answer' to store the rune result
	if s != ""{
		// Convert the first character of the string to a rune and store it in 'answer'
		answer = rune(s[0])
	}else{
		// If the string is empty, return the rune value 0 as the default case
		return 0
	}
	// Return the rune value stored in 'answer'
	return answer
}


/*
Get Last Rune
	Instructions
Create a Go function that takes a non-empty string as input and returns the last character of the string as a rune.
*/
func GetLastRune(s string) rune {
	// Initialize a variable 'answer' to store the rune result
	var answer rune
	// Convert the last character of the string to a rune and store it in 'answer'
	// This is achieved by accessing the character at position (length of s - 1)
	answer = rune(s[len(s)-1])
	// Return the rune value stored in 'answer'
	return answer
}


/*
Nth Rune
	Instructions
Create a Go function that takes a non-empty string as its first argument and an index as its second argument. The function should return the rune at the specified index in the string.
*/
func NRune(s string, i int) rune {
	// Initialize a variable 'answer' to store the rune result
	var answer rune
	// Convert the character at the i-th index of string 's' to a rune
	// and store it in 'answer'
	answer= rune(s[i])
	// Return the rune value stored in 'answer'
	return answer
}


/*
Str Lenght
	Instructions
Create a Go function that accepts a string and returns two integers. The first integer represents the number of runes in the string, and the second integer represents the byte length of the string. The function should handle UTF-8 encoding for all characters.
*/
func StrLength(s string) []int {
	// Initialize counters for rune and byte counts
	runeCount := 0
	byteCount := 0
	// Initialize an empty list 'answer' to store results
	var answer []int
	// Loop through each character in the string 's'
	for _, char := range s{
		// Increment the rune count for each character encountered
		runeCount++
		// Increment the byte count by the byte length of the current character
		//Increment byteCount by the length in bytes of the string representation of 'char'
		byteCount += len(string(char))
	}
	// Append the total rune count to the 'answer' list
	answer = append(answer, runeCount)
	// Append the total byte count to the 'answer' list
	answer = append(answer, byteCount)
	// Return the list 'answer' which contains the rune and byte counts
	return answer
}

/*
Str Reverse
	Instructions
Create a Go function that takes a string and returns the reversed version of that string.
*/
func StrReverse(s string) string {
	// Initialize an empty string 'answer' to store the reversed string
	var answer string
	// Loop from the end of the string 's' to the beginning
	for i:= len(s) -1; i>=0; i--{
		// Append the character at index i of string 's' to 'answer'
		answer += string(s[i])
	}
	// Return the reversed string stored in 'answer'
	return answer
}

/*
Is Lower?
	Instructions
Create a Go function that takes a string as its parameter. The function should return true if the string contains only lowercase characters, and false otherwise.
*/
func IsLower(s string) bool {
	// Loop through each character in the string 's'
	for _,char := range s{
		// Check if the character is not a lowercase letter
		if char < 'a' || char > 'z'{
			// If the character is not between 'a' and 'z', return false
			return false
		}
	}
	// If all characters are lowercase letters, return true
	return true
}

/*
Is Numeric?
	Instructions
Create a Go function that takes a string as its parameter. The function should return true if the string contains only numeric characters, and false otherwise.*/
func IsNumeric(s string) bool {
	// Loop through each character in the string 's'
	for _, num:= range s{
		// Check if the character is not a numeric digit
		if num < '0' || num > '9'{
			// If the character is not between '0' and '9', return false
			return false
		}
	}
	// If all characters are numeric digits, return true
	return true
}


/*
TO UPPER CASE
	Instructions
Create a Go function that takes a string as its parameter and returns a new string with each letter capitalized.
*/
func ToUpperCase(s string) string{
	// Initialize an empty string 'answer' to store the converted string
	var answer string
	 // Loop through each character in the string 's'
	for _, char := range s{
		// Check if the character is a lowercase letter
		if char >= 'a' && char <= 'z'{
			// Convert the lowercase letter to uppercase
			// The conversion is done by subtracting 'a' from 'char' to normalize (0-25),
			// taking modulo 26 as a safety measure (although not necessary here),
			// and then adding 'A' to shift to the uppercase range
			answer+= string((char-'a')%26 + 'A')
		}else{
			// If the character is not a lowercase letter, add it unchanged to 'answer'
			answer += string(char)
		}
	}
	// Return the converted string stored in 'answer'
	return answer
}


/*
To Capital Case
	Instructions
Create a Go function that takes a string as its parameter. The function should capitalize the first letter of each word while converting the rest of the word to lowercase.
In this task a word is defined as a sequence of alphanumeric characters.
*/
func ToCapitalCase(s string) string {
	// Initialize an empty string 'answer' to store the converted string
	var answer string
	// Initialize a boolean 'capitalize' to true to handle the first character as a capital letter
	capitalize := true
	// Loop through each character in the string 's'
	for _,char:= range s{
		// Check if the character is a non-alphanumeric (neither a letter nor a digit)
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9'){
			// Set 'capitalize' to true to capitalize the next alphabetic character
			capitalize = true
			// Add the non-alphanumeric character to 'answer'
			answer+= string(char)
		}else{
			// If 'capitalize' is true, process for capitalization
			if capitalize{
				// Check if 'char' is a lowercase letter
				if char >='a' && char <= 'z'{
					// Convert 'char' to an uppercase letter
					char= char-'a'+'A'
				}
				// Add the potentially modified 'char' to 'answer'
				answer+= string(char)
				// Set 'capitalize' to false after capitalizing
				capitalize= false

			} else{
				// If not capitalizing, ensure 'char' is in lowercase
				if char >= 'A' && char <= 'Z'{
					// Convert 'char' to a lowercase letter
					char += 'a'-'A'
				}
				// Add the lowercase or unchanged 'char' to 'answer'
				answer+= string(char)
			}
		}
	}
	// Return the converted string stored in 'answer'
	return answer
}


/*
Str Concat With
	Instructions
Create a Go function that takes a slice of strings and a separator string as its parameters. The function should return a new string by concatenating all the strings in the slice, with each string separated by the provided separator.
*/
func StrConcatWith(strs []string, sep string) string {
	// Initialize an empty string 'answer' to store the concatenated result
	var answer string
	// Loop through each string in the list 'strs' with its index 'i'
	for i, str:= range strs{
		// Check if the current string is the last in the list
		if i == len(strs)-1{
			// If it is the last string, append it to 'answer' without the separator
			answer+=str
		}else{
			// If it is not the last string, append the string and the separator to 'answer'
			answer+=str+sep
		}
	}
	// Return the concatenated string stored in 'answer'
	return answer
}

/*
Split on Whitespaces
	Instructions
Create a Go function that takes a string as its parameter. The function should split the string into words and store them in a string slice. Words are separated by spaces, tabs, and newlines.
*/
func SplitWhitespaces(s string) []string {
	// Initialize an empty list 'answer' to store the resulting substrings
	var answer []string
	// Initialize 'start' to 0 to mark the beginning of a new substring
	start := 0
	// Loop through each character in the string 's' with its index 'i'
	for i,char := range s{
		// Check if the current character is a whitespace (space, newline, or tab)
		if char == ' ' || char == '\n' || char == '\t' {
			// Append the substring from 'start' to 'i' (non-inclusive) to 'answer'
			answer= append(answer, s[start:i])
			// Update 'start' to 'i+1' to begin a new substring after the current whitespace
			start=i+1
		}
	}
	// After the loop, check if 'start' is not the end of the string
	if start != len(s){
		// Append the last substring from 'start' to the end of the string to 'answer'
		answer = append(answer, s[start:])
	}
	// Return the list of substrings stored in 'answer'
	return answer
}

/*
Str Split By
	Instructions
Create a Go function that takes a string s and a separator sep as parameters. The function should split the string s into substrings using the separator sep and return a slice of strings containing the resulting substrings.
*/
func StrSplitBy(s, sep string) []string {
	// Initialize an empty list 'answer' to store the resulting substrings
	answer:= []string{}
	// Initialize 'start' to 0 to mark the beginning of a new substring
	start := 0
	// Loop through the string 's' with index 'i' up to the position where separator can last occur
	for i:=0; i<= len(s)-len(sep);i++{
		// Check if the substring from 'i' to 'i+length of sep' equals 'sep'
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep{
			// If 'start' is not equal to 'i', append the substring from 'start' to 'i' to 'answer'
			if start != i {
				answer = append(answer, s[start:i])
			}
			// Update 'start' to skip the separator, 'i+length of sep'
			start = i+len(sep)
			// Increment 'i' by length of sep - 1 to skip checking characters part of current separator
			i+= len(sep)-1
		}
	}
	// After the loop, check if 'start' is less than the length of s
	if start < len(s){
		// Append the remaining substring from 'start' to the end of the string to 'answer'
		answer = append(answer, s[start:])
	}
	// Return the list of substrings stored in 'answer'
	return answer
}


/*
Substring Index
	Instructions
Create a Go function that behaves like the Index function. It takes two strings as input: s and toFind. The function should find the index of the first occurrence of toFind in s and return that index as an integer. If toFind is not present in s, the function should return -1.
*/
func SubstrIndex(s string, toFind string) int {
	// Loop through the string 's' from index 0 to the position where 'toFind' can last occur
	for i:= 0; i<= len(s)-len(toFind);i++{
		// Check if the substring from 'i' to 'i+length of toFind' equals 'toFind'
		if s[i:i+len(toFind)] == toFind{
			// If the substring is found, return the current index 'i'
			return i
		}
	}
	// If the loop completes without finding the substring, return -1
	return -1
}

/*
Str Compare
	Instructions
Create a Go function that mimics the behavior of the Compare function. It takes two strings, a and b, as input and returns an integer.
The function should compare the two strings and return:
	0 if the strings are equal,
	-1 if a comes before b in lexicographic order,
	1 if a comes after b in lexicographic order.
*/
func StrCompare(a, b string) int {
	// Loop through each character position 'i' as long as 'i' is less than the length of both strings
	for i:= 0; i < len(a) && i< len(b);i++{
		// Compare characters at position 'i' in both strings
		if a[i] < b[i]{
			// If character in 'a' is less, return -1 indicating 'a' is lexicographically smaller
			return -1
		}else if a[i]> b[i]{
			// If character in 'a' is greater, return 1 indicating 'a' is lexicographically larger
			return 1
		}
	}
	// After comparing all corresponding characters, check the lengths of the strings
	switch{
		case len(a) < len(b):
			// If string 'a' is shorter than string 'b', return -1
			return -1
		case len(a) > len(b):
			// If string 'a' is longer than string 'b', return 1
			return 1
	}
	// If all corresponding characters are equal and the lengths are the same, return 0
	return 0
}

/*
Number to Base
	Instructions
Create a Go function that takes an integer n and a string base as parameters. The function should return the integer n represented in the specified base as a string.
Here are the validity rules for the base:
	The base must contain at least 2 unique characters.
	The base should not contain the characters + or -.
	The function should handle negative numbers as well (see examples in the usage section). If the base is not valid, the function should return "NV" (Not Valid).
*/
func NbrBase(n int, base string) string {
	// Calculate the number of characters in 'base' as base_nr
	base_nr := len(base)
	// Initialize an empty string 'based_nr' to store the result
	based_nr := ""
	// Initialize 'neg' to 1 to handle positive numbers initially
	neg := 1
	// If 'n' is negative, adjust 'neg' to -1 and make 'n' positive
	if n < 0 {
		neg = -1
		n *= -1 //Multiply 'n' by -1 to make it positive
	}
	// Validate the base; it must have at least 2 characters and contain no '+' or '-'
	if base_nr < 2 {
		return "NV"
	}
	// Check each character in 'base' for invalid characters or duplicates
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' { // If character at position 'i' in base is '+' or '-':
			return "NV"
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] { // If character at position 'i' in base is equal to character at position 'j':
				return "NV"
			}
		}
	}
	// Convert number 'n' to the specified base
	for n > 0 {
		// Compute the remainder of 'n' divided by 'base_nr'
		modulo := n % base_nr
		// Get the character from 'base' corresponding to the remainder
		nr := base[modulo]
		// Prepend this character to 'based_nr'
		based_nr = string(rune(nr)) + based_nr
		// Divide 'n' by 'base_nr' to reduce it for the next iteration
		n = n / base_nr
	}
	// If the number was originally negative, prepend '-' to the result
	if neg == -1 {
		based_nr = "-" + based_nr
	}
	// Return the converted base number as a string
	return based_nr
}

/*
Convert Any to Decimal
	Instructions
Create a Go function that takes two parameters:
	s: a numeric string in a specific base.
	base: a string containing unique characters representing the available digits in that base.
	The function should return the integer value of s in the given base. If the base is not valid, it returns 0.
	Here are the validity rules for the base:
	The base must contain at least 2 unique characters.
	The base should not contain the characters + or -.
	The function only works with valid string numbers that consist of elements present in the base. It does not handle negative numbers.
Allowed packages:
	math
	strings
*/
func ConvertAnyToDec(s string, base string) int {
	// Validate the base string; it must contain at least two characters and no '+' or '-' symbols
	if len(base) < 2 || strings.ContainsAny(base, "+-") {
		return 0
	}
	// Create a map to store each character of the base string and its corresponding integer value
	valueMap := make(map[rune]int)
	// Populate the value map with indices of characters in the base string
	for i, char := range base {
		valueMap[char] = i
	}
	// Initialize 'answer' to 0 to accumulate the decimal value
	answer := 0
	// Convert the input string s to a decimal number using the base string
	for i := 0; i < len(s); i++ {
		// Get the value of the character at position 'i' in string s from the value map
		// Check if the character is valid (exists in the value map)
		value, is := valueMap[rune(s[i])]
		if !is {
			return 0 // Indicate an invalid character in s with a return value of 0
		}
		// Calculate the contribution of the current digit to the final number
		// by multiplying the character's base value by base's length raised to the appropriate power
		answer += value * int(math.Pow(float64(len(base)), float64(len(s)-i-1)))
	}
	// Return the computed decimal value
	return answer
}

/*
Convert Any to Any
	Instructions
Create a Go function that takes three string arguments: nbr representing a numeric value in a specific base baseFrom, and baseTo representing the desired base for the returned value. The function should convert the nbr from baseFrom to baseTo and return the result as a string.
Allowed packages:
	math
	strings
*/
func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	// Check for any invalid characters in 'nbr' that are not present in 'baseFrom'
	for _, char := range nbr {
		if !strings.ContainsRune(baseFrom, char) {
			return "Invalid character in the number"
		}
	}
	// Define a nested function to convert a number from any base to decimal
	convToDec := func(s, base string) int {
		// Create a map to store each character of the base and its corresponding integer value
		valueMap := make(map[rune]int)
		for i, char := range base {
			valueMap[char] = i // Set valueMap at key 'char' to 'i'
		}
		// Initialize 'answer' to 0 to accumulate the decimal value
		answer := 0
		for _, char := range s {
			// Retrieve the integer value of 'char' from the valueMap
			value, exists := valueMap[char]
			if !exists {
				return -1 // Character not found in base, invalid input
			}
			// Update 'answer' for each digit processed
			answer = answer*len(base) + value
		}
		return answer // Return the decimal value
	}
	// Define another nested function to convert a decimal number to any base
	convFromDec := func(n int, base string) string {
		// Handle zero specially
		if n == 0 {
			return string(base[0])
		}
		// Initialize an empty list of runes 'answer' to build the result in the target base
		answer := []rune{}
		length := len(base)
		for n > 0 {
			// Prepend the character corresponding to the remainder when n is divided by length of base
			answer = append([]rune{rune(base[n%length])}, answer...)
			// Reduce 'n' by dividing it by length of base
			n = n / length
		}
		return string(answer) // Convert list of runes to string and return
	}
	// Convert 'nbr' from 'baseFrom' to decimal
	decimal := convToDec(nbr, baseFrom)
	if decimal == -1 {
		return "Invalid number for the given base" // Input number doesn't properly conform to 'baseFrom'
	}
	// Convert the decimal value to 'baseTo'
	return convFromDec(decimal, baseTo)
}

func main(){
	fmt.Println(GetFirstRune("kood/Jõhvi"))
	fmt.Println(GetLastRune("kood/Jõhvi"))
	fmt.Println(NRune("kood/Jõhvi", 3))
	fmt.Println(StrLength("Hello World!"))
	fmt.Println(StrReverse("Hello World!"))
	fmt.Println(IsLower("hello"))
	fmt.Println(IsNumeric("010203"))
	fmt.Println(ToUpperCase("Hello! How are you?"))
	fmt.Println(ToCapitalCase("Hello! How are you? How+are+things+4you?"))
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(StrConcatWith(toConcat, ":"))
	fmt.Println(SplitWhitespaces("Hello how are you?"))
	s := "HelloHAhowHAareHAyou?"
	fmt.Println(StrSplitBy(s, "HA"))
	fmt.Println(SubstrIndex("Hello!", "l"))
	fmt.Println(StrCompare("Hello!", "Hello!"))
	fmt.Println(NbrBase(125, "0123456789"))
	fmt.Println(ConvertAnyToAny("101011", "01", "0123456789"))
}
