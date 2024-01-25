package phone_number_normalizer

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	table := []struct{ in, out string }{
		{in: "1234567890", out: "1234567890"},
		{in: "123 456 7891", out: "1234567891"},
		{in: "(123) 456 7892", out: "1234567892"},
		{in: "(123) 456-7893", out: "1234567893"},
		{in: "123-456-7894", out: "1234567894"},
		{in: "123-456-7890", out: "1234567890"},
		{in: "1234567892", out: "1234567892"},
		{in: "(123)456-7892", out: "1234567892"},
	}
	for _, testCase := range table {
		if res := Normalize(testCase.in); testCase.out != res {
			t.Errorf("Normalize(\"%s\") = \"%s\"", testCase.in, res)
		}
	}
}
