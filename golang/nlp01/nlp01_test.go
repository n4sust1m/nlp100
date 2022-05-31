package nlp01

import "testing"

func Test_pickEvenChar(t *testing.T) {
	value := pickEvenChar("パタトクカシーー")
	if value != "パトカー" {
		t.Errorf("failure: expect \"パトカー\", but return %v", value)
	}
}
