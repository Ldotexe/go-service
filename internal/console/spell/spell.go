package spell

import (
	"fmt"

	"homework-4/internal/console/command"
	"homework-4/internal/console/errors"
)

const newLineSymbol = rune(10)

type Command struct {
}

func New() command.Runner {
	return &Command{}
}

func (c *Command) Run(args []string) error {
	commandName := args[0]

	if len(args) != 2 {
		return errors.WrongArgsNumError(commandName)
	}

	word := args[1]
	fmt.Print(spell(word))
	return nil
}

func spell(word string) string {
	res := make([]rune, len(word)*2)
	for i, c := range word {
		res[i*2] = c
		res[i*2+1] = ' '
	}
	res[len(res)-1] = newLineSymbol
	return string(res)
}
