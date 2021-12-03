package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

type Command int

const (
	Unknown Command = iota
	Up
	Down
	Forward
)

func NewCommand(s string) (cmd Command, n int, err error) {
	s, n, err = split(s)
	if err != nil {
		return
	}
	switch s {
	case "up":
		cmd = Up
	case "down":
		cmd = Down
	case "forward":
		cmd = Forward
	default:
		err = fmt.Errorf("Unknown command: %s", s)
	}
	return
}

func split(s string) (cmd string, n int, err error) {
	ss := strings.Split(s, " ")
	if len(ss) != 2 {
		err = fmt.Errorf("line must have exactly 2 fields: %s", s)
		return
	}
	cmd = ss[0]
	n, err = strconv.Atoi(ss[1])
	return
}
