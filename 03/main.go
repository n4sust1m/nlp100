package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := []rune("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")

	var count_list []int
	count := 0
	for _, c := range str {
		if string(c) == " " {
			count_list = append(count_list, count)
			count = 0
			continue
		} else if unicode.IsUpper(c) || unicode.IsLower(c) {
			count++
		}
	}

	fmt.Println(count_list)

}
