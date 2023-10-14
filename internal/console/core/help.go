package console

import (
	"fmt"

	"homework-4/internal/console/command"
	"homework-4/internal/console/consolefmt"
	"homework-4/internal/console/delete"
	"homework-4/internal/console/get"
	"homework-4/internal/console/post"
	"homework-4/internal/console/put"
	"homework-4/internal/console/spell"
)

var (
	helpFirstPart = "Console is a tool for doing couple of things\n\nUsage:\n\n\t\tconsole <command> [arguments]\n\nThe commands are:\n\n"
)

func addCommand(commands map[string]command.Command, runner command.Runner) {
	cmd := runner.Add()
	commands[cmd.Name] = *cmd
}

func initCommands() map[string]command.Command {
	m := make(map[string]command.Command)

	addCommand(m, &spell.Command{})
	addCommand(m, &consolefmt.Command{})
	addCommand(m, &get.Command{})
	addCommand(m, &delete.Command{})
	addCommand(m, &post.Command{})
	addCommand(m, &put.Command{})

	return m
}

func Help(commands map[string]command.Command) error {
	longest := findLongest(commands)

	fmt.Print(helpFirstPart)

	for _, com := range commands {
		fmt.Printf("\t\t%s%s\t%s\n", com.Name, NSpaces(longest-len(com.Name)), com.Description)
	}

	fmt.Printf("\n")
	return nil
}

func findLongest(commands map[string]command.Command) int {
	longest := 0
	for _, com := range commands {
		if longest < len(com.Name) {
			longest = len(com.Name)
		}
	}
	return longest
}

func NSpaces(n int) string {
	spaces := ""
	for i := 0; i < n; i++ {
		spaces += " "
	}
	return spaces
}
