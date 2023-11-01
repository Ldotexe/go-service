package spell

import (
	"fmt"

	"homework-6/internal/console/command"
	"homework-6/internal/console/errors"
)

const newLineSymbol = rune(10)

type Command struct {
}

func New() command.Runner {
	return &Command{}
}

func (c *Command) Run(args []string) error {
	if len(args) != 1 {
		return errors.ErrWrongArgsNum
	}

	word := args[0]
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
