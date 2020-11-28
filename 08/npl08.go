package nlp08

const (
	_a = int('a')
	_z = int('z')
)

func crypt(str string) string {
	chars := []rune(str)
	result := []rune("")
	for i := range chars {
		char := int(chars[i])
		if _a <= char && char <= _z {
			char = 219 - char
		}
		result = append(result, rune(char))
	}
	return string(result)
}

func decrypt(str string) string {
	chars := []rune(str)
	result := []rune("")
	for i := range chars {
		char := int(chars[i])
		if _a <= 219-char && 219-char <= _z {
			char = 219 - char
		}
		result = append(result, rune(char))
	}
	return string(result)
}
