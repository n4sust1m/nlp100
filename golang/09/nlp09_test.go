package nlp09

import "testing"

func Test_typoglycemia(t *testing.T) {
	result := typoglycemia("I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind .")

	t.Logf("result: %+v", result)
}
