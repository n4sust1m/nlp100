package nlp08

import "testing"

func Test_crypt(t *testing.T) {
	result := crypt("Ok, Google! Tell me tommorow's weather.")
	t.Logf("%+v", result)
	unresult := decrypt(result)
	t.Logf("%+v", unresult)
}
