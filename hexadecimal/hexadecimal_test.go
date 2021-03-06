package hexadecimal

import (
	"testing"
)

var testCases = []struct {
	input           string
	expectedDecimal int64
	expectErr       bool
}{
	{"1", 1, false},
	{"10", 16, false},
	{"2d", 45, false},
	{"cfcfcf", 13619151, false},
	{"CFCFCF", 13619151, false},
	{"peanut", 0, true},
	{"2cg134", 0, true},
}

func TestParseHex(t *testing.T) {
	for _, test := range testCases {
		actualDecimal, actualErr := ParseHex(test.input)
		if actualDecimal != test.expectedDecimal {
			t.Fatalf("ParseHex(%s): expected result[%d], actual[%d]",
				test.input, test.expectedDecimal, actualDecimal)
		}
		// if we expect an error and there isn't one
		if test.expectErr && actualErr == nil {
			t.Errorf("ParseHex(%s): expected an error, but error is nil", test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectErr && actualErr != nil {
			t.Errorf("ParseHex(%s): expected no error, but error is: %s", test.input, actualErr)
		}
	}
}

func BenchmarkParseHex(b *testing.B) {
	b.StopTimer()

	for _, test := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ParseHex(test.input)
		}

		b.StopTimer()
	}
}
