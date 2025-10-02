package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct{
		input string
		expected []string
	}{
		{input: "z f 34 ", expected: []string{"z", "f", "34"}},
		{input: "iii", expected: []string{"iii"}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {t.Errorf("length mismatch: expected %d but got %d (input %q)", len(c.expected), len(actual), c.input)}

		mismatch := false
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("word %d mismatch: expected %q got %q (input %q)", i, c.expected[i], actual[i], c.input)
				mismatch = true
			}
		}
		if !mismatch {
            t.Logf("case %v passed with %+v", c.input, actual)
		}
	}
}

