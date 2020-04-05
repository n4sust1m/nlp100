package main

import (
	"fmt"
	"strings"
)

func main  () {
	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."

	splitted_str := strings.Split(str, " ")

	m := map[string]string{}

	for i, s := range splitted_str {
		switch i {
		case 0, 4, 5, 6, 7, 8, 14, 15, 18:
			m[string([]rune(s)[0])] = s
		default:
			m[string([]rune(s)[0:2])] = s
		}
	}
	fmt.Print(m)
}
