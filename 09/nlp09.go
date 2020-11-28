package nlp09

import (
	"math/rand"
	"regexp"
	"time"
)

func Typoglycemia(str string) string {
	re := regexp.MustCompile(`(\ )+`)
	words := re.Split(str, -1)
	result := ""
	for i, v := range words {
		if i != 0 {
			result += " "
		}
		if len(v) >= 5 {
			result += _shuffleMiddleChars(v)
		} else {
			result += v
		}
	}
	return result
}

func _shuffleMiddleChars(str string) string {
	rn := []rune(str)

	f := rn[0]
	rn = rn[1:]

	l := rn[len(rn)-1]
	rn = rn[:len(rn)-1]

	result := []rune{f}
	index := len(rn)
	for index > 0 {
		rand.Seed(time.Now().UnixNano())
		rdm := rand.Intn(index)
		result = append(result, rn[rdm])
		rn = append(rn[:rdm], rn[rdm+1:]...)
		index--
	}
	result = append(result, l)
	return string(result)
}
