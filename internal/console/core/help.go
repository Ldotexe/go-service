package console

import (
	"fmt"
	"strings"

	"homework-6/internal/console/command"
)

var (
	helpFirstPart = "Console is a tool for doing couple of things\n\nUsage:\n\n\tconsole <command> [arguments]\n\nThe commands are:\n\n"
)

func Help(commands map[string]command.Command) error {
	longest := findLongest(commands)

	fmt.Print(helpFirstPart)

	for _, com := range commands {
		fmt.Printf("\t%s%s   %s\n", com.Name, strings.Repeat(" ", longest-len(com.Name)), com.Description)
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
