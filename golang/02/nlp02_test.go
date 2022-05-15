package nlp02

import "testing"

func Test_mix(t *testing.T) {
	if mix("パトカー", "タクシー") != "パタトクカシーー" {
		t.Errorf("error")
	}
}
