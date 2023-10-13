package consolefmt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode"

	"homework-4/internal/console/command"
	"homework-4/internal/console/errors"
)

type Command struct {
}

func (c *Command) Add() *command.Command {
	return command.NewCommand(
		"consolefmt", "inserts a tab before each paragraph and puts a dot at the end of sentences for *.txt file",
		c.Run,
	)
}

func (c *Command) Run(args []string) error {
	commandName := args[1]

	if len(args) != 3 {
		return errors.NewErrWrongArgsNum(commandName)
	}

	filename := args[2]
	if !regexp.MustCompile(`*.txt`).MatchString(filename) {
		return errors.ErrWrongFormat
	}

	return format(filename)
}

func format(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	in := read(r)

	out := process(in)

	w, err := os.OpenFile(filename)
	if err != nil {
		return err
	}
	write(w, out)

	return nil
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
	res := ""
	if unicode.IsUpper(rune(line[0])) {
		res += "\t"
	}
	for _, c := range line {
		if unicode.IsUpper(c) {

		}
	}
}

func process(lines []string) []string {
	out := make([]string, 0)
	for _, line := range lines {
		out = append(out, processLine(line))
	}
	return out
}

func write(res *WordMap, in []string) {
	for procRes := range in {
		for i, word := range res.words {
			res.mp[word] += procRes.keywords[i]
		}
	}

	for _, word := range res.words {
		res.num += res.mp[word]
	}
}

func showResults(wm WordMap) {
	for _, word := range wm.words {
		fmt.Printf("%s: %d\n", word, wm.mp[word])
	}
	fmt.Printf("всего: %d\n", wm.num)
}
