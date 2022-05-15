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

/*

func main() {
	str1 := "paraparaparadise"
	str2 := "paragraph"

	x := c_bi_gram(str1)
	y := c_bi_gram(str2)

	var sum []string
	var diff []string
	var prod []string

	var is_se_in_prod bool = false

	for x_key := range x {
		// 和集合
		sum = append(sum, x_key)

		if _, ok := y[x_key]; ok {
			// 積集合
			prod = append(prod, x_key)
			if x_key == "se" {
				is_se_in_prod = true
			}
		} else {
			// 差集合
			diff = append(diff, x_key)
		}
	}
	for y_key := range y {
		if _, ok := x[y_key]; !ok {
			// 和集合
			sum = append(sum, y_key)
		}
	}
	fmt.Println("和集合", sum)
	fmt.Println("差集合", diff)
	fmt.Println("積集合", prod)
	fmt.Println("\"se\" in 積集合? ==", is_se_in_prod)
}

func c_bi_gram(s string) map[string]int {
	splitted_str := strings.Split(s, "")
	length := len(splitted_str)

	m := map[string]int{}
	for i := 0; length-1 > i; i++ {
		this_key := splitted_str[i] + splitted_str[i+1]
		if _, ok := m[this_key]; ok {
			m[this_key]++
		} else {
			m[this_key] = 1
		}
	}
	return m
}
*/
