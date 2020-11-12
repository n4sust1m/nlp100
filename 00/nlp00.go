package nlp00

func reverse(str string) string {
	tmp := []rune(str)
	for i := 0; i < len(str)/2; i = i + 1 {
		tmp[i], tmp[len(str)-1-i] = tmp[len(str)-1-i], tmp[i]
	}
	return string(tmp)
}
