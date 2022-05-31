// http://www.cl.ecei.tohoku.ac.jp/nlp100/#sec06
package nlp06

import (
	"regexp"

	mapset "github.com/deckarep/golang-set"
)

func make2gram(str string) []string {
	chars := regexp.MustCompile(``).Split(str, -1)
	tmp := []string{}
	for i := 0; i < len(chars); i = i + 2 {
		c1 := chars[i]
		c2 := ""
		if i+1 < len(chars) {
			c2 = chars[i+1]
		}
		tmp = append(tmp, c1+c2)
	}
	return tmp
}

func sum(x []string, y []string) mapset.Set {
	result := mapset.NewSet()
	for _, v := range x {
		result.Add(v)
	}
	for _, v := range y {
		result.Add(v)
	}
	return result
}

func diff(x []string, y []string) mapset.Set {
	_x := mapset.NewSet()
	for _, v := range x {
		_x.Add(v)
	}
	_y := mapset.NewSet()
	for _, v := range y {
		_y.Add(v)
	}

	result := mapset.NewSet()
	for v := range _x.Iter() {
		if !_y.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func prod(x []string, y []string) mapset.Set {
	_x := mapset.NewSet()
	for _, v := range x {
		_x.Add(v)
	}
	_y := mapset.NewSet()
	for _, v := range y {
		_y.Add(v)
	}

	result := mapset.NewSet()
	for v := range _x.Iter() {
		if _y.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func contain(x []string, v string) bool {
	for _, _v := range x {
		if v == _v {
			return true
		}
	}
	return false
}
