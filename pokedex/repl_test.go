package main

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, desc, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\n*** %#v ***\n Expected %#v, got %#v", desc, want, got)
	}
}

func TestStuff(t *testing.T) {
	assertEqual(t, "Spaces front and back and middle",
		cleanInput("   	hello	world   "),
		[]string{"hello", "world"})
	assertEqual(t, "Spaces front and back and middle, amount",
		len(cleanInput("   	hello	world   ")),
		2)
	assertEqual(t, "mixed case words",
		cleanInput("Charmander Balbasaur PIKACHU"),
		[]string{"charmander", "balbasaur", "pikachu"})
	assertEqual(t, "empty input",
		cleanInput(""),
		[]string{""})
}

func TestCommandExplore(t *testing.T) {
	cfg := &config{}

	err := commandExplore(cfg)
	assertEqual(t, "explore with no argumnents",
		err.Error(),
		"you must provide one a location name")

	// Test case: too many arguments
	err = commandExplore(cfg, "pastoria", "extra-arg")
	assertEqual(t, "explore with too many arguments",
		err.Error(),
		"you must provide one a location name")
}
