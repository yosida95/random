package random

import (
	"testing"
)

func TestReader(t *testing.T) {
	r, err := String(UPPER | LOWER | DIGITS)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	rr, _, err := r.ReadRune()
	if err == nil {
		println(string(rr))
	} else {
		t.Log(err)
		t.Fail()
	}

}
