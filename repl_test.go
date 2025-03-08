package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "With UpperCase ",
			expected: []string{"with", "uppercase"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected len %v, got %v", len(c.expected), len(actual))
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Expected %v, got %v", c.expected, actual)
			}
		}
	}
}
