package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	answer := strings.Join(strings.Fields(lowered), " ")
	return strings.Split(answer, " ")
}
