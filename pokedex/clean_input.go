package cleanInput

import "strings"

func cleanInput(text string) []string {
	return strings.Split(text, " ")

}
