package console

import (
	ers "errors"

	"homework-6/internal/console/commands"
	"homework-6/internal/console/errors"
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
			if ers.Is(err, errors.ErrWrongArgsNum) {
				return errors.WrongArgsNumError(name)
			}
			return err
		}

		return nil
	}
	return errors.UnknownCommandError(name)
}
