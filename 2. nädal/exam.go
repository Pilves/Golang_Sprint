package main

import (
	"fmt"
	"strconv"
)

func toLower(s string) string {
	var answer string

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			answer += string(char - 'A' + 'a')
		} else {
			answer += string(char)
		}
	}

	return answer
}

func isUpper(s string) bool {
	for _, char := range s {
		if (char < 'A' || char > 'Z') && char != ' ' {
			return false
		}
	}
	return true
}

func rot13(s string) string {
	var answer string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			answer += string(int(char-'a'+13)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			answer += string(int(char-'A'+13)%26 + 'A')
		} else {
			answer += string(char)
		}
	}
	return answer
}

func fizzArray(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, strconv.Itoa(i))
		}
	}
	return answer
}

func uniqueAcc(n []int) int {
	var answer int
	uniques := make(map[int]bool)
	for _, num := range n {
		if !uniques[num] {
			uniques[num] = true
			answer += num
		}
	}
	return answer
}

func fromBinary(s string) int {
	var result int
	for _, digit := range s {
		if digit == '1' {
			result = result*2 + 1
		} else if digit == '0' {
			result *= 2
		} else {
			return -1
		}
	}
	return result
}

func CombN(n int) []string {
	if n <= 0 {
		return nil
	}
	result := []string{}
	stack := make([]int, n)
	for i := range stack {
		stack[i] = i
	}
	for {
		comb := ""
		for _, v := range stack {
			comb += fmt.Sprintf("%d", v)
		}
		result = append(result, comb)
		i := n - 1
		for i >= 0 && stack[i] == 9-(n-1-i) {
			i--
		}
		if i < 0 {
			break
		}
		stack[i]++
		for j := i + 1; j < n; j++ {
			stack[j] = stack[j-1] + 1
		}
	}
	return result
}

func main() {
	fmt.Println(toLower("THIS is A TeST W0rd"))
	fmt.Println(isUpper("HELLO WORLD"))
	fmt.Println(rot13("Abc!123 Nop!123"))
	fmt.Println(fizzArray(25))
	intList := []int{1, 2, 3, 4, 2, 5, 6, 3, 7, 8, 9, 5, 10}
	fmt.Println(uniqueAcc(intList))
	binaryStr := "11110100011000111000111"
	fmt.Println(fromBinary(binaryStr))
	fmt.Println(CombN(1))
}
