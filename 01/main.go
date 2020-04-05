package main

import "fmt"

func main() {
	str := []rune("パタトクカシー")
	for i := 0; i < len(str); i = i + 2 {
		fmt.Print(string(str[i]))
	}
	fmt.Print("\n")
}
