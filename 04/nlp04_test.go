package nlp04

import (
	"reflect"
	"testing"
)

func TestPickInitial2Chars(t *testing.T) {

	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	opt := Option{[]int{1, 5, 6, 7, 8, 9, 15, 16, 19}}
	result := pickInitial2Chars(str, opt)
	if !reflect.DeepEqual(
		result,
		map[int]string{
			0:  "H",
			1:  "He",
			2:  "Li",
			3:  "Be",
			4:  "B",
			5:  "C",
			6:  "N",
			7:  "O",
			8:  "F",
			9:  "Ne",
			10: "Na",
			11: "Mi",
			12: "Al",
			13: "Si",
			14: "P",
			15: "S",
			16: "Cl",
			17: "Ar",
			18: "K",
			19: "Ca",
		}) {
		t.Errorf("ErrorL output in unexepected value -> %v", result)
	} else {
		t.Logf("successfully: output -> %v", result)
	}
}
