package nlp03

import (
	"reflect"
	"testing"
)

func Test_countCharactorsInSentense(t *testing.T) {
	result := countCharactorsInSentense("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	expect := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}
	if !reflect.DeepEqual(
		result,
		expect,
	) {
		t.Errorf("Error: not expected output -> %v", result)
	}
}
