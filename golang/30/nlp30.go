package nlp30

import (
	"errors"
	"os"
	"strings"
)

type Morpheme struct {
	surface string // 表層形
	base    string // 基本形
	pos     string // 品詞
	pos1    string // 品詞分類
}

func LoadMecab() (*[]Morpheme, error) {
	bytes, err := os.ReadFile("../tmp/neko.txt.mecab")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")

	r := make([]Morpheme, len(lines))
	for index, line := range lines {
		if line == "EOS" {
			continue
		}

		s := strings.Split(line, "	")
		r[index].surface = s[0]
		if len(s) == 1 {
			continue
		}
		m := strings.Split(s[1], ",")
		if len(m) < 6 {
			return nil, errors.New("invalid format")
		}
		r[index].base = m[6]
		r[index].pos = m[0]
		r[index].pos1 = m[1]

		index++
	}

	return &r, nil
}
