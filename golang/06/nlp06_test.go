package nlp06

import (
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set"
)

const (
	str1 = "paraparaparadise"
	str2 = "paragraph"
)

func Test_make2Gram(t *testing.T) {
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

	expectedSet := mapset.NewSet()
	for _, v := range []string{"pa", "ra", "di", "se", "gr", "ap", "h"} {
		expectedSet.Add(v)
	}

	result := sum(result1, result2)
	if result.Equal(expectedSet) {
		t.Logf("verify: %+v", result)
	} else {
		t.Errorf("Invalid Result: %+v", result)
	}
}

func TestDiff(t *testing.T) {
	result1 := make2gram(str1)
	result2 := make2gram(str2)
	expectedSet := mapset.NewSet()
	for _, v := range []string{"di", "se"} {
		expectedSet.Add(v)
	}

	result := diff(result1, result2)
	if !expectedSet.Equal(result) {
		t.Errorf("invalid result: %+v, expect %+v", result, expectedSet)
	} else {
		t.Logf("verify: %+v", result)
	}
}

func TestProd(t *testing.T) {
	result1 := make2gram(str1)
	result2 := make2gram(str2)
	expectedSet := mapset.NewSet()
	for _, v := range []string{"pa", "ra"} {
		expectedSet.Add(v)
	}
	result := prod(result1, result2)
	if !result.Equal(expectedSet) {
		t.Errorf("invalid result: %+v, expect %+v", result, expectedSet)
	} else {
		t.Logf("verify: %+v", result)
	}
}

func TestContain(t *testing.T) {
	if !contain(make2gram(str1), "se") {
		t.Errorf("invalid result")
	}
	if contain(make2gram(str2), "se") {
		t.Errorf("invalid result")
	}
}
