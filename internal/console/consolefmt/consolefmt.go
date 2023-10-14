package consolefmt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"homework-4/internal/console/command"
	"homework-4/internal/console/errors"
)

type Command struct {
}

func New() command.Runner {
	return &Command{}
}

func (c *Command) Run(args []string) error {
	if len(args) != 1 {
		return errors.ErrWrongArgsNum
	}

	filename := args[0]
	if !strings.HasSuffix(filename, ".txt") {
		return errors.ErrWrongFormatFilename
	}

	return format(filename)
}

func format(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	in := read(f)

	f, err = os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	out := process(in)

	return write(f, out)
}

func read(r io.Reader) []string {
	s := bufio.NewScanner(r)

	out := make([]string, 0)
	for s.Scan() {
		out = append(out, s.Text())
	}

	return out
}

func processLine(line string) string {
	if !strings.HasPrefix(line, "\t") && !strings.HasPrefix(line, "    ") {
		line = "\t" + line
	}
	if !strings.HasSuffix(line, ".") {
		line += "."
	}
	return line
}

func process(lines []string) []string {
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		out = append(out, processLine(line))
	}
	return out
}

func write(w io.Writer, in []string) error {
	for _, str := range in {
		_, err := fmt.Fprintln(w, str)
		if err != nil {
			return err
		}
	}
	return nil
}
