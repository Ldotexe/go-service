package post

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
		"post",
		"runs the post command with the ID, name and points specified in the arguments",
		c.Run,
	)
}

func (c *Command) Run(args []string) error {
	commandName := args[1]

	if len(args) != 5 {
		return errors.NewErrWrongArgsNum(commandName)
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.ErrWrongFormatId
	}
	name := args[3]
	points, err := strconv.Atoi(args[4])
	if err != nil {
		return errors.ErrWrongFormatPoints
	}
	return post(id, name, points)
}

func post(id int, name string, points int) error {
	cmd := exec.Command(
		"curl", "-X", "POST", "localhost:9000/student", "-d",
		fmt.Sprintf("{\"id\":%d,\"name\":\"%s\",\"points\":%d}", id, name, points), "-i",
	)

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
