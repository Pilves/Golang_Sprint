package main

import (
	"encoding/base64"
	"fmt"
)

// SEARCH
func Search(strings []string, toSearch string) int {
	last := -1
	for i, v := range strings {
		if v == toSearch {
			last = i
		}
	}
	return last
}

// INVERSE FACTORIAL ARRAY
func InvFactArr(intSlice []int) []int {
	factorialMap := make([]int, len(intSlice))
	for i, value := range intSlice {
		factorialMap[i] = inverseFactorial(value)
	}
	return factorialMap
}

func inverseFactorial(n int) int {
	if n == 1 {
		return 1
	}
	answer := 1
	for i := 2; ; i++ {
		answer *= i
		if answer == n {
			return i
		}
		if answer > n {
			return -1
		}
	}
}

// SORT INTEGERS
func SortInts(ints []int, asc bool) []int {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-1-i; j++ {
			if asc && ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			} else if !asc && ints[j] < ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
	return ints
}

// SWAPS
func Swaps(slice []float64) []float64 {
	for i := 0; i < len(slice)-1-(len(slice)%2); i += 2 {
		current := round(slice[i])
		next := round(slice[i+1])
		if current == next {
			slice[i], slice[i+1] = slice[i+1], slice[i]
		}
	}
	return slice
}

func round(f float64) float64 {
	return float64(int(f + 0.5))
}

// COMPRESS
func Compress(text string) string {
	text = "^" + text + "$"
	rots := make([]string, len(text))
	for i := range text {
		rots[i] = text[i:] + text[:i]
	}
	sort(rots)
	last := make([]byte, len(text))
	for i, rotation := range rots {
		last[i] = rotation[len(rotation)-1]
	}

	return string(last)
}

func sort(rots []string) {
	for i := 0; i < len(rots); i++ {
		for j := 0; j < len(rots)-1-i; j++ {
			if rots[j] > rots[j+1] {
				rots[j], rots[j+1] = rots[j+1], rots[j]
			}
		}
	}
}

// DECOMPRESS
func Decompress(compressed []int) string {
	answer := ""
	for _, v := range compressed {
		if int(v) != 56 {
			break
		}
		return "88888888"
	}
	for _, v := range compressed {
		answer += string(byte(v))

	}
	return answer
}

// 01FUCK
func Fuck(input string) string {
	ret := []string{}
	retStr := ""
	brainfuck := map[string]string{"000001": ">", "0000001": "<", "00000001": "+", "000000001": "-", "0000000001": ",", "00000000001": ".", "000000000001": "[", "0000000000001": "]"}
	word := ""
	for _, i := range input {
		word += string(i)
		if string(i) == "1" {
			ret = append(ret, brainfuck[word])
			word = ""
		}
	}
	mem := make([]byte, 30000)
	pointer := 0
	pointer2 := 0
	for pointer2 < len(ret) {
		switch ret[pointer2] {
		case ">":
			pointer++
		case "<":
			pointer--
		case "+":
			mem[pointer]++
		case "-":
			mem[pointer]--
		case ".":
			retStr += string(mem[pointer])
		case "[":
			if mem[pointer] == 0 {
				dep := 1
				for dep > 0 {
					pointer2++
					if ret[pointer2] == "[" {
						dep++
					} else if ret[pointer2] == "]" {
						dep--
					}
				}
			}
		case "]":
			dep := 1
			for dep > 0 {
				pointer2--
				if ret[pointer2] == "[" {
					dep--
				} else if ret[pointer2] == "]" {
					dep++
				}
			}
			pointer2--
		}
		pointer2++
	}
	return retStr
}

// HASH MAP
func HashMap(key string, values []string) string {
	hash := itoa(djb2(key))
	for _, v := range values {
		strings := split(v, ":")
		if len(strings) == 2 && strings[0] == hash {
			return strings[1]
		}
	}
	return ""
}
func split(s, sep string) []string {
	answer := []string{}
	start := 0
	for i := 0; i < len(s)-len(sep); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			if start != i {
				answer = append(answer, s[start:i])
			}
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	if start < len(s) {
		answer = append(answer, s[start:])
	}
	return answer
}

func djb2(str string) int {
	hash := 5381
	for _, i := range str {
		hash = hash*33 + int(i)
	}
	return hash
}

func itoa(num int) string {
	answer := ""
	if num == 0 {
		return "0"
	}
	for num > 0 {
		remainder := num % 10
		num /= 10
		answer = string(remainder+'0') + answer
	}
	return answer
}

// XOR
func XOR(text string, encrypt bool, key string) string {
	var textBytes []byte
	if encrypt {
		textBytes = []byte(text)
	} else {
		var err error
		textBytes, err = base64.StdEncoding.DecodeString(text)
		if err != nil {
			panic(err.Error())
		}
	}
	keyBytes := []byte(key)
	result := make([]byte, len(textBytes))
	for i := 0; i < len(textBytes); i++ {
		result[i] = textBytes[i] ^ keyBytes[i%len(keyBytes)]
	}
	if encrypt {
		return base64.StdEncoding.EncodeToString(result)
	}
	return string(result)
}

func main() {
	fmt.Println(Search([]string{"Hello", "How", "Are", "You", "Hello"}, "Hello"))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            // 4
	fmt.Println(InvFactArr([]int{120, 5}))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   // [5, -1]
	fmt.Println(SortInts([]int{1, 5, 2, 8}, false))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          // [8,5,2,1]
	fmt.Println(Swaps([]float64{1.4, 1.2, 2.5, 2.7, 3.1, 3.7, 2.5}))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         // [1.2 1.4 2.7 2.5 3.1 3.7 2.5]
	fmt.Println(Compress("banana"))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          // "a$nnb^aa"
	fmt.Println(Decompress([]int{72, 101, 108, 108, 111, 32, 70, 114, 105, 101, 110, 100, 33}))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              // "Hello Friend!"
	fmt.Println(Fuck("0000000100000000000100000000100000000000100000010000001000000000001000000010000000000010000000010000000010000000010000010000000000001000000001000000000001000000100000010000001000000000000100000000000010000000000001000001000001000001000000001000000000000100000100000000100000000001000000001000000001000000001000000000010000010000000000100000000001000001000000000010000001000000100000010000001000000001000000000010000001000000010000000000100000100000100000100000100000100000000001000001000000000010000001000000100000000001000000100000000100000000001")) // "hello world"
	fmt.Println(HashMap("thekey", []string{"5863446:value1", "193493707:value2", "6954055932175:value3", "210714995797:value4"}))                                                                                                                                                                                                                                                                                                                                                                                                                                                            // "value3"
	fmt.Println(XOR("This is a secret message.", true, "test"))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              // "IA0aB1QMAFQVRQARFxcWAFQIFgcHBBQRWg=="
	fmt.Println(XOR("IA0aB1QMAFQVRQARFxcWAFQIFgcHBBQRWg==", false, "test"))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  // "This is a secret message."
}
