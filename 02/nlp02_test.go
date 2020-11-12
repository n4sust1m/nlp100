package nlp02

import "testing"

func TestMix(t *testing.T) {
	if mix("パトカー", "タクシー") != "パタトクカシーー" {
		t.Errorf("error")
	}
}
