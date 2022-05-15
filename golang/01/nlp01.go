package nlp01

func pickEvenChar(arg string) string {
	str := []rune(arg)
	tmp := []rune("")
	for i := 0; i < len(str); i = i + 2 {
		tmp = append(tmp, str[i])
	}
	return string(tmp)
}
