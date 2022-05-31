package nlp02

func mix(str1 string, str2 string) string {
	str1Tmp := []rune(str1)
	str2Tmp := []rune(str2)

	tmp := []rune("")
	for i := range str1Tmp {
		tmp = append(tmp, str1Tmp[i], str2Tmp[i])
	}
	return string(tmp)
}
