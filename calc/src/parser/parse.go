package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse parses an expression in format 'a op b'
func Parse(expr string) (string, int, int, error) {
	args := strings.Split(expr, " ")
	if len(args) < 3 {
		return "", 0, 0, fmt.Errorf("not enough args")
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return "", 0, 0, err
	}
	b, err := strconv.Atoi(args[2])
	if err != nil {
		return "", 0, 0, err
	}
	return args[1], a, b, nil
}
