package re

import "testing"

var RgxObjct Rgx = CreateRe()

func TestCaseSensitiveRgxMatch(t *testing.T) {
	cases := []struct {
		input    string
		rgxPttrn string
		expected bool
	}{
		{
			input: "Hello", rgxPttrn: "Hello", expected: true,
		},
	}

	for _, cs := range cases {
		actual, _ := RgxObjct.In(cs.rgxPttrn, cs.input, true)
		if actual != cs.expected {
			t.Errorf("Actual: %v , Expected: %v", actual, cs.expected)
		}
	}
}

func TestCaseInSensitiveRgxMatch(t *testing.T) {
	cases := []struct {
		input    string
		rgxPttrn string
		expected bool
	}{
		{
			input: "Hello", rgxPttrn: "hello", expected: false,
		},
	}

	for _, cs := range cases {
		actual, _ := RgxObjct.In(cs.rgxPttrn, cs.input, false)
		if actual != cs.expected {
			t.Errorf("Actual: %v , Expected: %v", actual, cs.expected)
		}
	}
}
