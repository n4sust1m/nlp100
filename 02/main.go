package main

import (
	"fmt"
)

func main() {
	str1 := []rune("パトカー")
	str2 := []rune("タクシー")
	for i := range str1 {
		fmt.Print(string(str1[i]))
		fmt.Print(string(str2[i]))
	}
	fmt.Print("\n")
}
