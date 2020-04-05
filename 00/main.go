package main

import "fmt"

func main() {
	str := "stressed"
	tmp := []rune(str)
	for i := 0; i < len(str)/2; i = i + 1 {
		tmp[i], tmp[len(str)-1-i] = tmp[len(str)-1-i], tmp[i]
	}
	fmt.Println(string(tmp))
}
