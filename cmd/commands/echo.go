package commands

import "strings"

func Echo(args []string) CommandOutput {
	if len(args) == 0 {
		return CommandOutput{
			0, "", nil,
		}
	}

	return CommandOutput{
		0, strings.Join(args, " "), nil,
	}
}
