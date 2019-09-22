package main

import "fmt"

func main() {
	var str = "パタトクカシーー"

	var rn = []rune(str)

	var result = ""

	var i = 0
	for i = 1; i < len(rn); i = i + 2 {
		result += string(rn[i])
	}

	fmt.Print(result)
}
