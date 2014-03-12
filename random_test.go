package random

import (
	"testing"
)

type asciiTestCase struct {
	flag  asciiflag
	check func(rune) bool
}

var (
	asciiTestCases = []asciiTestCase{
		asciiTestCase{
			flag:  UPPER,
			check: isUpperAscii,
		},
		asciiTestCase{
			flag:  LOWER,
			check: isLowerAscii,
		},
		asciiTestCase{
			flag:  DIGITS,
			check: isDigits,
		},
		asciiTestCase{
			flag: UPPER | LOWER,
			check: func(r rune) bool {
				return isUpperAscii(r) || isLowerAscii(r)
			},
		},
		asciiTestCase{
			flag: UPPER | DIGITS,
			check: func(r rune) bool {
				return isUpperAscii(r) || isDigits(r)
			},
		},
		asciiTestCase{
			flag: LOWER | DIGITS,
			check: func(r rune) bool {
				return isLowerAscii(r) || isDigits(r)
			},
		},
		asciiTestCase{
			flag: UPPER | LOWER | DIGITS,
			check: func(r rune) bool {
				return isUpperAscii(r) || isLowerAscii(r) || isDigits(r)
			},
		},
	}
)

func TestReader(t *testing.T) {
	r, err := newReader(26, 65)
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

		if !isUpperAscii(rr) {
			t.Logf("returned rune is out of range, %d", rr)
			t.Fail()
			continue
		}
	}
}

func TestAscii(t *testing.T) {
	for _, testCase := range asciiTestCases {
		r, err := Ascii(testCase.flag)
		if err == nil {
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

				if !testCase.check(rr) {
					t.Logf("returned rune is out of range, %d", rr)
					t.Fail()
					continue
				}
			}
		} else {
			t.Log(err)
			t.Fail()
		}
	}
}

func isUpperAscii(r rune) bool {
	return 65 <= r && r <= 90
}

func isLowerAscii(r rune) bool {
	return 97 <= r && r <= 122
}

func isDigits(r rune) bool {
	return 48 <= r && r <= 57
}
