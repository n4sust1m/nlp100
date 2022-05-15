//str := "I am an NLPer"

package nlp05

import (
	"reflect"
	"testing"
)

func Test_make2Gram(t *testing.T) {
	str := "I am an NLPer"
	result := make2Gram(str)

	if !reflect.DeepEqual(result.charactor2gram, []string{"Ia", "ma", "nN", "LP", "er"}) {
		t.Errorf("Error: invalid char2gram")
	} else {
		t.Logf("verify: %#v", result.charactor2gram)
	}
	if !reflect.DeepEqual(result.word2gram, []string{"I am", "an NLPer"}) {
		t.Errorf("Error: invalid word2gram")
	} else {
		t.Logf("verify: %#v", result.word2gram)
	}
}
