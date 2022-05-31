package nlp03

import (
	"regexp"
)

func countCharactorsInSentense(arg string) []int {
	matchPattern := regexp.MustCompile(`(\W)+`)
	chars := matchPattern.Split(arg, -1)

	var countList []int
	for _, c := range chars {
		if len(c) != 0 {
			countList = append(countList, len(c))
		}
	}
	return countList
}
