package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	got := cleanInput("	hello	world	")
	want := []string{"hello", "world"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
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
