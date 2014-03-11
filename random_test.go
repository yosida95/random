package random

import (
	"testing"
)

func TestReader(t *testing.T) {
	r, err := newReader(10, 0)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for i := 0; i < 255; i++ {
		rr, i, err := r.ReadRune()
		if err != nil {
			t.Log(err)
			t.Fail()
			continue
		}

		if i > 1 {
			t.Logf("length of rune must be 1 but returned one is %d", i)
			t.Fail()
			continue
		}

		if rr < 0 || rr >= 10 {
			t.Logf("returned rune is out of range, %d", rr)
			t.Fail()
			continue
		}
	}
}
