package update

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"

	"homework-6/internal/console/command"
	"homework-6/internal/console/errors"
)

type Command struct {
}

func New() command.Runner {
	return &Command{}
}

func (c *Command) Run(args []string) error {
	if len(args) != 3 {
		return errors.ErrWrongArgsNum
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.ErrWrongFormatId
	}
	name := args[1]
	points, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.ErrWrongFormatPoints
	}
	return update(id, name, points)
}

func update(id int, name string, points int) error {
	cmd := exec.Command(
		"curl", "-X", "PUT", "localhost:9000/student", "-d",
		fmt.Sprintf("{\"id\":%d,\"name\":\"%s\",\"points\":%d}", id, name, points), "-i",
	)

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
