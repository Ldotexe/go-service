package errors

import (
	"errors"
	"fmt"
)

var EndOfErr = "Run 'console help' for usage."

var (
	ErrWrongFormatFilename = errors.New("wrong format for argument *.txt")
	ErrWrongFormatId       = errors.New("wrong format for argument ID")
	ErrWrongFormatPoints   = errors.New("wrong format for argument points")
	ErrCommandAlreadyExist = errors.New("command already exist")
)

func UnknownCommandError(command string) error {
	return errors.New(fmt.Sprintf("console %s: unknon command\n%s", command, EndOfErr))
}

func WrongArgsNumError(command string) error {
	return errors.New(fmt.Sprintf("console %s: wrong number of arguments\n%s", command, EndOfErr))
}
