package get

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"

	"homework-4/internal/console/command"
	"homework-4/internal/console/errors"
)

type Command struct {
}

func (c *Command) Add() *command.Command {
	return command.NewCommand(
		"get",
		"runs get command",
		c.Run,
	)
}

func (c *Command) Run(args []string) error {
	commandName := args[1]

	if len(args) != 3 {
		return errors.NewErrWrongArgsNum(commandName)
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.ErrWrongFormatId
	}
	return getID(id)
}

func getID(id int) error {
	cmd := exec.Command("curl", "-X", "GET", fmt.Sprintf("localhost:9000/student/%d", id), "-i")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	fmt.Printf("%s\n", string(data))
	return err
}
