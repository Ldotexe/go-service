package console

import (
	"homework-4/internal/console/errors"
)

func Run(args []string) error {
	commands := initCommands()
	if len(args) == 0 || args[0] == "help" {
		return Help(commands)
	}
	name := args[0]
	com, ok := commands[name]
	if ok {
		return com.Runner.Run(args)
	}
	return errors.NewErrUnknownCommand(name)
}
