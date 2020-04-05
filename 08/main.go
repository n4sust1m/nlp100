package main

import (
	"fmt"
)

const diff rune = 'z' - 'a'
func isLower (c rune) bool {
	if c >= 'a' && 'z' >= c {
		return true
	}
	return false
}
func isUpper (c rune) bool {
	if c >= 'A' && 'Z' >= c {
		return true
	}
	return false
}

func main () {
	str := "こんにちは、Hello World 世界。"
	fmt.Println("msg:", str)
	str = to_cipher(str)
	fmt.Println("ciphered msg:", str)
	str = from_cipher(str)
	fmt.Println("deciph msg:", str)
}

func to_cipher(str string) string {
	converted := []rune{}
	for _, c := range str {
		if isUpper(c) {
			c = 219 - c
			for !isUpper(c) {
				c = diff + c 
			}
		} else if isLower(c) {
			c = 219 - c
			for !isLower(c) {
				c = diff + c 
			}
		}
		converted = append(converted, c)
	}
	return string(converted)
}

func from_cipher(str string) string {
	converted := []rune{}
	for _, c := range str {
		if isUpper(c) {
			c = 219 + c
			for !isUpper(c) {
				c = c - diff 
			}
		} else if isLower(c) {
			c = 219 + c
			for !isLower(c) {
				c = c - diff 
			}
		}
		converted = append(converted, c)
	}
	return string(converted)
}