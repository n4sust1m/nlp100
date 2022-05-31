package nlp05

import (
	"regexp"
	"strings"
)

type SplitPatterns struct {
	word      *regexp.Regexp
	charactor *regexp.Regexp
}

type Results struct {
	word2gram      []string
	charactor2gram []string
}

func make2Gram(str string) Results {
	pats := SplitPatterns{
		word:      regexp.MustCompile(`(\W)+`),
		charactor: regexp.MustCompile(``),
	}
	words := pats.word.Split(str, -1)
	chars := pats.charactor.Split(str, -1)

	result := Results{
		word2gram:      []string{},
		charactor2gram: []string{},
	}

	// make word 2-gram sequence
	for i := 0; i < len(words); i = i + 2 {
		result.word2gram = append(result.word2gram, words[i]+" "+words[i+1])
	}

	// make charactor 2-gram sequence
	var tmp = ""
	for i := range chars {
		if 0 == strings.Compare(" ", chars[i]) {
			continue
		}
		tmp += chars[i]
		if len(tmp) >= 2 {
			result.charactor2gram = append(result.charactor2gram, tmp)
			tmp = ""
		}
	}

	return result
}
