package nlp30

import (
	"reflect"
	"testing"
)

func Test_LoadMecab(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		m, err := LoadMecab()
		if err != nil {
			t.Errorf("err: %T", err)
		}
		if reflect.TypeOf(*m).String() == "[]Morpheme" {
			t.Error("unexpected type")
		}
	})
}
