package command

type Command struct {
	Name        string
	Description string
	Run         func(args []string) error
}

func NewCommand(name string, description string, run func(args []string) error) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Run:         run,
	}
}

type Usage interface {
	Add() *Command
	Run(args []string) error
}
