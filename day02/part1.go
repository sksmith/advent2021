package day02

import (
	"strconv"
	"strings"
)

func Move(dirs []string) int {
	var x, depth int

	for _, s := range dirs {
		dir, dst := ParseCmd(s)

		switch dir {
		case "forward":
			x += dst
		case "down":
			depth += dst
		case "up":
			depth -= dst
		}
	}

	return x * depth
}

func MoveAdvanced(dirs []string) int {
	var aim, x, depth int

	for _, s := range dirs {
		dir, dst := ParseCmd(s)

		switch dir {
		case "forward":
			x += dst
			depth += aim * dst
		case "down":
			aim += dst
		case "up":
			aim -= dst
		}
	}

	return x * depth
}

func ParseCmd(s string) (string, int) {
	cmd := strings.Split(s, " ")
	dir := cmd[0]
	dst, err := strconv.Atoi(cmd[1])

	if err != nil {
		panic(err)
	}
	return dir, dst
}
