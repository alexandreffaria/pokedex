package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := make(map[string]cliCommand)

	commands["help"] = cliCommand{
		name:        "help",
		description: "shows help message",
		callback:    commandHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "exit program",
		callback:    commandExit,
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			commands["exit"].callback()
		} else if input == "help" {
			commands["help"].callback()
		}
		fmt.Print("pokedex > ")
	}

}

func commandHelp() error {
	fmt.Println("Command HELP")
	return nil
}

func commandExit() error {
	fmt.Println("Command EXIT")
	return nil
}
