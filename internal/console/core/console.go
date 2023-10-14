package console

import (
	"homework-4/internal/console/commands"
	"homework-4/internal/console/errors"
)

func Run(args []string) error {
	initCommands, err := commands.InitCommands()
	if err != nil {
		return err
	}
	if len(args) == 0 || args[0] == "help" {
		return Help(initCommands)
	}
	name := args[0]
	com, ok := initCommands[name]
	if ok {
		return com.Runner.Run(args)
	}
	return errors.UnknownCommandError(name)
}
