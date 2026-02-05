package cleanInput

import (
	"strings"
)

func cleanInput(text string) []string {
	answer := strings.Join(strings.Fields(text), " ")
	return strings.Split(answer, " ")
}
