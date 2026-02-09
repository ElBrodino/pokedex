package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())

		cmd, ok := commands[text[0]]
		if len(text) == 0 {
			continue
		}

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Println("Error:", err)
		}
		//fmt.Println("Your command was:", text[0])

	}
}
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	answer := strings.Join(strings.Fields(lowered), " ")
	return strings.Split(answer, " ")
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
}
