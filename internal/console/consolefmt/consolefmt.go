package consolefmt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"homework-4/internal/console/command"
	"homework-4/internal/console/errors"
)

type Command struct {
}

func (c *Command) Add() *command.Command {
	return command.NewCommand(
		"consolefmt", "inserts a tab before each paragraph and puts a dot at the end of sentences for *.txt file",
		&Command{},
	)
}

func (c *Command) Run(args []string) error {
	commandName := args[1]

	if len(args) != 3 {
		return errors.NewErrWrongArgsNum(commandName)
	}

	filename := args[2]
	ok, err := regexp.MatchString(`.*\.txt`, filename)
	if err != nil {
		return err
	}
	if !ok {
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

func wordIsFirst(word string) bool {
	ok, err := regexp.MatchString(`^[A-Z][a-z]*`, word)
	if err != nil {
		return false
	}
	return ok
}

func wordIsLast(word string) bool {
	ok, err := regexp.MatchString(`.*\.`, word)
	if err != nil {
		return false
	}
	return ok
}

func processLine(line string) string {
	res := ""
	s := strings.Split(line, " ")
	for i, word := range s {
		if wordIsFirst(word) {
			if i == 0 {
				res += "\t"
			} else {
				if wordIsLast(s[i-1]) {
					res += " "
				} else {
					res += ". "
				}
			}
		} else {
			if i != 0 {
				res += " "
			}
		}
		res += word

	}
	return res
}

func process(lines []string) []string {
	out := make([]string, 0)
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
