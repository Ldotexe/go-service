package errors

import (
	"errors"
	"fmt"
)

var EndOfErr = "Run 'console help' for usage."

var (
	ErrWrongFormatFilename = errors.New("wrong format for argument *.txt")
	ErrWrongFormatId       = errors.New("wrong format for argument ID")
)

func NewErrUnknownCommand(command string) error {
	return errors.New(fmt.Sprintf("console %s: unknon command\n%s", command, EndOfErr))
}

func NewErrWrongArgsNum(command string) error {
	return errors.New(fmt.Sprintf("console %s: wrong number of arguments\n%s", command, EndOfErr))
}
