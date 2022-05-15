package nlp07

import (
	"strings"
	"testing"
)

func Test_template(t *testing.T) {
	result := template("12", "気温", "22.4")
	if strings.Compare(result, "12時の気温は22.4") != 0 {
		t.Errorf("Invalid result: %+v, expected %+v", result, "12時の気温は22.4")
	} else {
		t.Logf("verify: %+v", result)
	}
}
