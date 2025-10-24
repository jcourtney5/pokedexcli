package main

import (
	"fmt"
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
			input:    "a  b  c",
			expected: []string{"a", "b", "c"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "the quick brown fox",
			expected: []string{"the", "quick", "brown", "fox"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := cleanInput(c.input)

			// Check the length of the actual slice against the expected slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if len(actual) != len(c.expected) {
				t.Errorf("cleanInput(%q) expected length %v, actual length %v", c.input, len(c.expected), len(actual))
			}

			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				// Check each word in the slice
				// if they don't match, use t.Errorf to print an error message
				// and fail the test
				if word != expectedWord {
					t.Errorf("cleanInput(%q) word number %v not equal expected %q, actual %q", c.input, i, expectedWord, word)
				}
			}
		})
	}

	//for _, c := range cases {
	//	actual := cleanInput(c.input)
	//
	//	// Check the length of the actual slice against the expected slice
	//	// if they don't match, use t.Errorf to print an error message
	//	// and fail the test
	//	if len(actual) != len(c.expected) {
	//		t.Errorf("cleanInput(%q) expected length %v, actual length %v", c.input, len(c.expected), len(actual))
	//	}
	//
	//	for i := range actual {
	//		word := actual[i]
	//		expectedWord := c.expected[i]
	//		// Check each word in the slice
	//		// if they don't match, use t.Errorf to print an error message
	//		// and fail the test
	//		if word != expectedWord {
	//			t.Errorf("cleanInput(%q) word number %v not equal expected %q, actual %q", c.input, i, expectedWord, word)
	//		}
	//	}
	//}
}
