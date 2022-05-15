package nlp04

import (
	"regexp"
)

type Option struct {
	index1char []int
}

func pickInitial2Chars(str string, option Option) map[int]string {
	splitPattern := regexp.MustCompile(`(\W)+`)
	chars := splitPattern.Split(str, -1)

	result := map[int]string{}

	for index, c := range chars {
		if len(c) == 0 {
			continue
		}
		var pickInitial int = indexOf(option.index1char, index+1)
		if pickInitial > -1 {
			result[index] = string([]rune(c)[0])
		} else {
			result[index] = string([]rune(c)[0:2])
		}
	}
	return result
}

func indexOf(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
