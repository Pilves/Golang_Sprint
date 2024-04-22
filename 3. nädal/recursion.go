package main

import(
	"fmt"
)
/*
Factorial Iterative
	Instructions
You're tasked with creating an iterative function that calculates the factorial of an integer passed as a parameter. Make sure to handle errors, returning 0 for non-possible values or overflows.
*/

func FactorialIterative(n int) int {
	 // Initialize the variable to store the factorial result
	answer := 1
	 // Check if the input is negative
	if n < 0 {
		// If it's negative, return 0 since factorial is not defined for negative numbers
		return 0
	}
	// Loop from 1 to n (inclusive)
	for i:= 1; i <= n; i++{
		 // Multiply the current answer by the loop counter
		answer *= i
	}
	// Return the factorial result
	return answer
}

/*
Factorial Recursive
	Instructions
You're tasked with creating a recursive function that calculates the factorial of an integer passed as a parameter. Make sure to handle errors, returning 0 for non-possible values or overflows.
*/

func FactorialRecursive(n int) int {
	// Base case: If n is negative, return 0 since factorial is not defined for negative numbers
	if n < 0{
		return 0
	}
	 // Base case: If n is 0, return 1 because the factorial of 0 is defined as 1
	if n == 0{
		return 1
	}
	// Recursive case: Multiply n by the factorial of (n-1) and return the result
	return n* FactorialRecursive(n-1)
}

/*
To the Power Iterative
	Instructions
You're tasked with creating an iterative function that calculates the power of an integer n to the given power. Handle negative powers by returning 0.
*/

func ToThePowerIterative(n int, power int) int {
	// Check if the power is negative
	if power < 0 {
		// If power is negative, return 0 since raising to a negative power is undefined
		return 0
	}
	 // Initialize the variable to store the result
	answer := 1
	// Check if the power is 0
	if power == 0{
		// If power is 0, return 1 since any number raised to the power of 0 is 1
		return 1
	}
	// Check if the power is 1
	if power == 1{
		// If power is 1, return n since any number raised to the power of 1 is itself
		return n
	}
	// Loop from 0 to power-1
	for i:= 0; i< power;i++{
		// Multiply the current answer by n
		answer *= n
	}
	// Return the final result
	return answer
}

/*
To the Power Recursive
	Instructions
You're tasked with creating an recursive function that calculates the power of an integer n to the given power. Handle negative powers by returning 0.
*/

func ToThePowerRecursive(n int, power int) int {
	// Check if the power is negative
	if power < 0{
		// If power is negative, return 0 since raising to a negative power is undefined
		return 0
	}
	// Base case: If power is 0, return 1 since any number raised to the power of 0 is 1
	if power == 0{
		return 1
	}
	// Recursive case: Multiply n by the result of (n^(power-1)) and return the result
	return n * ToThePowerRecursive(n, power-1)
}

/*
Nth Fibonacci
	Instructions
You're tasked with creating a recursive function that returns the value at a given position index in the Fibonacci sequence. The first value is at index 0.
The Fibonacci sequence starts with 0 and 1, and each subsequent value is the sum of the two preceding values: 0, 1, 1, 2, 3, etc.
Negative indices should be handled by returning -1.
*/

func NthFibonacci(index int) int {
	// Check if the index is negative
	if index < 0{
		// If index is negative, return -1 indicating an invalid input
		return -1
	}
	// Base case: If index is 0, return 0 since the 0th Fibonacci number is 0
	 if index == 0{
		 return 0
	// Base case: If index is 1, return 1 since the 1st Fibonacci number is 1
	}else if index == 1{
		return 1
	}
	// Recursive case: Return the sum of the (index-1)th and (index-2)th Fibonacci numbers
	return NthFibonacci(index-1) + NthFibonacci(index-2)
}

/*
Is Prime?
	Instructions
You're tasked with creating a function that determines whether the integer passed as a parameter is a prime number. The function should return true if the number is prime and false if it's not.
The function must be optimized to avoid time-outs during testing.
We consider that only positive numbers can be prime numbers, and 1 is not considered a prime number.
*/

func IsPrime(n int) bool{
	// Check if n is less than or equal to 1
	if n <= 1{
		// If n is less than or equal to 1, return false since prime numbers are greater than 1
		return false
	}
	// Loop from i = 2 to the square root of n
	for i := 2; i*i <= n; i++{
		// Check if n is divisible by i
		if n % i == 0{
			// If n is divisible by i, return false indicating that n is not a prime number
			return false
		}
	}
	// If the loop completes without finding any divisors of n, return true indicating that n is a prime number
	return true
}

/*
Next Prime
	Instructions
You're tasked with creating a function that returns the first prime number that is equal to or greater than the integer passed as a parameter. The function should return this prime number.
The function must be optimized to avoid time-outs during testing.
We consider that only positive numbers can be prime numbers.
*/

func NextPrime(n int) int {
	// Check if n is less than or equal to 1
	if n <= 1{
		// If n is less than or equal to 1, return 2 since 2 is the smallest prime number
		return 2
	}
	// Start a loop from num = n
	for num := n; ; num++{
		// Check if num is a prime number using the helper function IsPrime **note** you need to rewrite this func to NextPrime file
		if IsPrime(num){
			// If num is a prime number, return num
			return num
		}
	}
}

/*
Sqrt
	Instructions
You're tasked with creating a function that returns the square root of the integer passed as a parameter, but only if that square root is a whole number. Otherwise, it should return 0.
*/

func Sqrt(n int) int {
	// Check if n is less than 2
	if n < 2{
		// If n is less than 2, return n since the square root of 0 or 1 is the number itself
		return n
	}
	// Loop from i = 2 to n/2
	for i:= 2; i <= n/2; i++{
		// Check if i*i is equal to n
		if i*i ==n {
			// If i*i is equal to n, return i since i is the square root of n
			return i
		}
	}
	// If the loop completes without finding the square root, return 0 indicating that n is not a perfect square
	return 0
}

func main(){
	fmt.Println(FactorialIterative(4))
	fmt.Println(FactorialRecursive(4))
	fmt.Println(ToThePowerIterative(4, 3))
	fmt.Println(ToThePowerRecursive(4, 3))
	fmt.Println(NthFibonacci(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
	fmt.Println(NextPrime(5))
	fmt.Println(NextPrime(4))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
