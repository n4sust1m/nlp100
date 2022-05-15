package nlp01

import "testing"

func TestPickEvenChar(t *testing.T) {
	value := pickEvenChar("パタトクカシーー")
	if value != "パトカー" {
		t.Errorf("failure: expect \"パトカー\", but return %v", value)
	}
}
