package main

import (
	"testing"	
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"HOW what why", []string{"how", "what", "why"}},
		{"Wy why", []string{"wy", "why"}},
	}

	for _, c := range tests {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Slice length mismatch")
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Word mismatch")
			}
		}
	}
}
