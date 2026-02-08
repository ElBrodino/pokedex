package main

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected %#v, got %#v", want, got)
	}
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		//name     string
		input    string
		expected []string
	}{
		{
			//name:     "spaces front and back",
			input:    "	hello	world	",
			expected: []string{"hello", "world"},
		},
		{
			//name:     "mixed case words",
			input:    "Charmander Balbasaur PIKACHU",
			expected: []string{"charmander", "balbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		assertEqual(t, actual, c.expected)
	}
}
