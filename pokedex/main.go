package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello, World!")

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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())

		cmd, ok := commands[text[0]]
		if !ok {
			cleanInput(text[0])
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Println("Error:", err)
		}
		//fmt.Println("Your command was:", text[0])

	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
