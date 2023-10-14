package command

type Command struct {
	Name        string
	Description string
	Runner      Runner
}

func NewCommand(name string, description string, run Runner) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Runner:      run,
	}
}

type Runner interface {
	Run(args []string) error
	Add() *Command
}
