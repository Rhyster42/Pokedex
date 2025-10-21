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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello THERE friend",
			expected: []string{"hello", "there", "friend"},
		},
		{
			input:    "May the foRce be WITH YOU",
			expected: []string{"may", "the", "force", "be", "with", "you"},
		},
		{
			input:    "    cash OR        credit     ",
			expected: []string{"cash", "or", "credit"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error: Returned slice wrong length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: %s does not match %s", word, expectedWord)
			}
		}
	}
}
