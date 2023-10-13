package spell

import (
	"fmt"

	"homework-4/internal/console/command"
	"homework-4/internal/console/errors"
)

type Command struct {
}

func (c *Command) Add() *command.Command {
	return command.NewCommand(
		"spell",
		"takes a word as input and displays all the letters of that word separated by a space to the console",
		c.Run,
	)
}

func (c *Command) Run(args []string) error {
	commandName := args[1]

	if len(args) != 3 {
		return errors.NewErrWrongArgsNum(commandName)
	}

	word := args[2]
	fmt.Print(spell(word))
	return nil
}

func spell(word string) string {
	res := ""
	for _, c := range word {
		res += fmt.Sprintf("%c ", c)
	}
	res += "\n"
	return res
}
