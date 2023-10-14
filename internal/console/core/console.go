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
		err = com.Runner.Run(args[1:])
		if err != nil {
			if err == errors.ErrWrongArgsNum {
				return errors.WrongArgsNumError(name)
			}
			return err
		}

		return nil
	}
	return errors.UnknownCommandError(name)
}
