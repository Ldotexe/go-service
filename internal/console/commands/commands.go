package commands

import (
	ers "errors"

	"homework-4/internal/console/command"
	"homework-4/internal/console/consolefmt"
	"homework-4/internal/console/delete"
	"homework-4/internal/console/errors"
	"homework-4/internal/console/get"
	"homework-4/internal/console/post"
	"homework-4/internal/console/spell"
	"homework-4/internal/console/update"
)

func addCommand(commands map[string]command.Command, name string, description string, runner command.Runner) error {
	_, ok := commands[name]
	if ok {
		return errors.ErrCommandAlreadyExist
	}
	commands[name] = *command.NewCommand(name, description, runner)
	return nil
}

func InitCommands() (map[string]command.Command, error) {
	m := make(map[string]command.Command)
	err := ers.Join(

		addCommand(
			m, "spell",
			"takes a word as input and displays all the letters of that word separated by a space to the console",
			spell.New(),
		),

		addCommand(
			m,
			"consolefmt", "inserts a tab before each paragraph and puts a dot at the end of sentences for *.txt file",
			consolefmt.New(),
		),

		addCommand(
			m, "get",
			"runs the get command with the ID specified in the argument", get.New(),
		),

		addCommand(
			m, "delete",
			"runs the delete command with the ID specified in the argument", delete.New(),
		),

		addCommand(
			m, "post",
			"runs the post command with the ID, name and points specified in the arguments", post.New(),
		),

		addCommand(
			m, "update",
			"runs the update command with the ID, name and points specified in the arguments", update.New(),
		),
	)

	return m, err
}
