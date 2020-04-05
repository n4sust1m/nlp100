package main

import (
	"fmt"
	"strings"
)

func word_bi_gram(s string) map[string]int {
	words := strings.Split(s, " ")
	length := len(words)

	m := map[string]int{}
	for i := 0; i < length-1; i++ {
		this_key := words[i] + " " + words[i+1]
		if _, ok := m[this_key]; ok {
			m[this_key]++
		} else {
			m[this_key] = 1
		}
	}
	return m
}

func charactor_bi_gram(s string) map[string]int {
	chars := strings.Split(s, "")
	m := map[string]int{}

	length := len(chars)
	for i := 0; i < length-1; i++ {
		var nextChar = chars[i+1]
		if chars[i] == " " {
			continue
		} else if chars[i+1] == " " {
			nextChar = chars[i+2]
		}
		this_key := chars[i] + nextChar
		if _, ok := m[this_key]; ok {
			m[this_key]++
		} else {
			m[this_key] = 1
		}
	}
	return m
}

func main() {
	str := "I am an NLPer"

	wb := word_bi_gram(str)
	fmt.Println("Words bi-gram", wb)

	cb := charactor_bi_gram(str)
	fmt.Println("Charactors bi-gram", cb)
}
