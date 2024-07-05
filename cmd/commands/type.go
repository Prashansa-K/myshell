package commands

import (
	"fmt"
	"strings"
)

const (
	SHELL_BUILTIN = " is a shell builtin"
	NOT_FOUND     = ": not found"
)

func Type(args []string) CommandOutput {
	token := strings.TrimSpace(args[0])

	for _, c := range validCommands {
		if token == c {
			return CommandOutput{
				0, fmt.Sprintf("%s%s\n", c, SHELL_BUILTIN), nil,
			}
		}
	}

	return CommandOutput{
		0, fmt.Sprintf("%s%s\n", token, NOT_FOUND), nil,
	}
}
