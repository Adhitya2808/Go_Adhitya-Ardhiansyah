package main

import "fmt"

func caesar(offset int, input string) string {
	hasil := ""

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			newChar := 'a' + ((char - 'a' + int32(offset)) % 26)
			hasil += string(newChar)
		} else {
			hasil += string(char)
		}
	}

	return hasil
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // def
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}