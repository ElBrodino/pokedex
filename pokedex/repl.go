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
	callback    func(*config) error
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl() {

	cfg := &config{
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cmd, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg); err != nil {
			fmt.Println("Error:", err)
		}

	}
}
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	answer := strings.Join(strings.Fields(lowered), " ")
	return strings.Split(answer, " ")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List 20 first locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List 20 previous maps",
			callback:    commandMapB,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
