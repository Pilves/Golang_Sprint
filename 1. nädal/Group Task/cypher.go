package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	toEncrypt, encoding, message, err := getInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ciphers := map[string]func(string) string{
		"Rot13":   encrypt_rot13,
		"Reverse": encrypt_reverse,
		"Base64":  encrypt_base64,
	}
	deciphers := map[string]func(string) string{
		"Rot13":   decrypt_rot13,
		"Reverse": decrypt_reverse,
		"Base64":  decrypt_base64,
	}

	cipherFunc := ciphers[encoding]
	decipherFunc := deciphers[encoding]
	if toEncrypt {
		fmt.Println("Result:", cipherFunc(message))
	} else {
		fmt.Println("Result:", decipherFunc(message))
	}
}

func getInput() (toEncrypt bool, encoding string, message string, err error) {
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println("Select operation (1: Encrypt, 2: Decrypt):")
	var operationInt int
	if _, err = fmt.Scan(&operationInt); err != nil {
		return
	}
	toEncrypt = operationInt == 1
	if operationInt != 1 && operationInt != 2 {
		err = fmt.Errorf("invalid operation selection")
		return
	}
	fmt.Println("Select cipher (1: Rot13, 2: Reverse 3: Base64):")
	var cipherInt int
	if _, err = fmt.Scan(&cipherInt); err != nil {
		return
	}
	switch cipherInt {
	case 1:
		encoding = "Rot13"
	case 2:
		encoding = "Reverse"
	case 3:
		encoding = "Base64"
	default:
		err = fmt.Errorf("invalid cipher selection")
		return
	}
	fmt.Println("Enter your message:")
	reader := bufio.NewReader(os.Stdin)
	inputMessage, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	message = strings.TrimSpace(inputMessage)
	return
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {

	var encrypted string
	for _, char := range s {
		switch {
		case char >= 'a' && char <= 'z':
			encrypted += string('a' + (char-'a'+13)%26)
		case char >= 'A' && char <= 'Z':
			encrypted += string('A' + (char-'A'+13)%26)
		case char >= '0' && char <= '9':
			encrypted += string('0' + (char-'0'+5)%10)
		default:
			encrypted += string(char)
		}
	}
	return encrypted

}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {
	reversed := ""
	for _, character := range s {
		switch {
		case (character >= 'a' && character <= 'z'):
			reversed += string('z' - (character - 'a'))
		case (character >= 'A' && character <= 'Z'):
			reversed += string('Z' - (character - 'A'))
		case (character >= '0' && character <= '9'):
			reversed += string('9' - (character - '0'))
		default:
			reversed += string(character)
		}
	}
	return reversed
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}

// Encrypt the message with base64
func encrypt_base64(s string) string {
	const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	binaryString := ""
	for _, char := range s {
		binaryString += fmt.Sprintf("%08b", char)
	}
	for len(binaryString)%6 != 0 {
		binaryString += "0"
	}
	var decimalNumberString string
	for i := 0; i < len(binaryString); i += 6 {
		end := i + 6
		sixBits := binaryString[i:end]
		decimalValue, err := strconv.ParseInt(sixBits, 2, 64)
		if err != nil {
			fmt.Println("Error converting binary to decimal:", err)
			continue
		}
		decimalNumberString += string(base64Chars[decimalValue])
	}
	return decimalNumberString
}

// Decrypt the message with base64
func decrypt_base64(encoded string) string {
	const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	reverseBase64Chars := make(map[byte]int)
	for index, char := range base64Chars {
		reverseBase64Chars[byte(char)] = index
	}
	binaryString := ""
	for i := 0; i < len(encoded); i++ {
		char := encoded[i]
		index := reverseBase64Chars[char]
		binaryString += fmt.Sprintf("%06b", index)
	}
	for len(binaryString)%8 != 0 {
		binaryString += "0"
	}
	var decodedString string
	for i := 0; i < len(binaryString); i += 8 {
		byteString := binaryString[i : i+8]
		asciiValue, err := strconv.ParseInt(byteString, 2, 64)
		if err != nil {
			fmt.Println("Error converting binary to ASCII:", err)
			continue
		}
		decodedString += string(asciiValue)
	}
	return decodedString
}
