package console

import (
	"homework-4/internal/console/errors"
)

func Run(args []string) error {
	commands := SetHelp()
	if len(args) == 1 || args[1] == "help" {
		return Help(commands)
	}
	name := args[1]
	com, ok := commands[name]
	if ok {
		return com.Run(args)
	}
	return errors.NewErrUnknownCommand(name)
}
