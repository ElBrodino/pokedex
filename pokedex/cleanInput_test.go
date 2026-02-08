package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "spaces front and back",
			input:    "	hello	world	",
			expected: []string{"hello", "world"},
		}, /*
			{
				name:     "mixed case words",
				input:    "Charmander Balbasaur PIKACHU",
				expected: []string{"charmander", "balbasaur", "pikachu"},
			},*/
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := cleanInput(c.input)

			if len(got) != len(c.expected) {
				t.Fatalf("len mismatch: got %v, expected %v", got, c.expected)
			}

			for i := range got {
				if got[i] != c.expected[i] {
					t.Fatalf("cleanInput(%q)[%d] = %q, expected %q",
						c.input, i, got[i], c.expected[i])
				}
			}
		})
	}
}

/*
cases := []struct {
	input	string
	expected []string
}{
	input:	"	hello	world	",
	expected: []string{"hello", "world",
	},
	// add more cases here
}
/*
for _, c := cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}
}
*/
