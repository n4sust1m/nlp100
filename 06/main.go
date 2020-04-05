package main

import (
	"fmt"
	"strings"
)

func main () {
	str1 := "paraparaparadise"
	str2 := "paragraph"

	x := c_bi_gram(str1)
	y := c_bi_gram(str2)
	
	var sum []string
	var diff []string
	var prod []string

	var is_se_in_prod bool = false

	for x_key := range x {
		// 和集合
		sum = append(sum, x_key)
		
		if _,ok := y[x_key]; ok {
			// 積集合
			prod = append(prod, x_key)
			if x_key == "se" {
				is_se_in_prod = true
			}
		} else {
			// 差集合
			diff = append(diff, x_key)
		}
	}
	for y_key := range y {
		if _, ok := x[y_key]; !ok {
			// 和集合
			sum = append(sum, y_key)
		}
	}
	fmt.Println("和集合", sum)
	fmt.Println("差集合", diff)
	fmt.Println("積集合", prod)
	fmt.Println("\"se\" in 積集合? ==", is_se_in_prod)
}

func c_bi_gram (s string) map[string]int {
	splitted_str := strings.Split(s, "")
	length := len(splitted_str)
	
	m := map[string]int{}
	for i := 0; length - 1 > i; i++ {
		this_key := splitted_str[i] + splitted_str[i+1]
		if _, ok := m[this_key]; ok {
			m[this_key]++
		}	else {
			m[this_key] = 1
		}
	}
	return m
}