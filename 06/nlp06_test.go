package nlp06

import (
	"reflect"
	"testing"
)

const (
	str1 = "paraparaparadise"
	str2 = "paragraph"
)

func TestMake2Gram(t *testing.T) {
	result1 := make2gram(str1)
	result2 := make2gram(str2)

	if !reflect.DeepEqual(result1, []string{"pa", "ra", "pa", "ra", "pa", "ra", "di", "se"}) {
		t.Errorf("invalid result: %#v", result1)
	} else {
		t.Logf("verify: %#v", result1)
	}
	if !reflect.DeepEqual(result2, []string{"pa", "ra", "gr", "ap", "h"}) {
		t.Errorf("invalid result: %#v", result2)
	} else {
		t.Logf("verify: %#v", result2)
	}
}

func TestSum(t *testing.T) {

	result1 := make2gram(str1)
	result2 := make2gram(str2)

	result := sum(result1, result2)
	if reflect.DeepEqual(result, []string{"pa", "ra", "di", "se", "gr", "ap", "h"}) {
		t.Logf("verify: %#v", result)
	} else {
		t.Errorf("Invalid Result: %#v", result)
	}
}
