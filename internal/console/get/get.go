package get

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"

	"homework-4/internal/console/command"
	"homework-4/internal/console/errors"
)

type Command struct {
}

func New() command.Runner {
	return &Command{}
}

func (c *Command) Run(args []string) error {
	commandName := args[0]

	if len(args) != 2 {
		return errors.WrongArgsNumError(commandName)
	}

	id, err := strconv.Atoi(args[1])
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

	data, err := io.ReadAll(stdout)
	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	fmt.Printf("%s\n", string(data))
	return err
}
