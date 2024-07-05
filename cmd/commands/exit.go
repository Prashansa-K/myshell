package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Exit(args []string) CommandOutput {
	if len(args) == 0 { // only exit passed, assuming it as exit 0
		os.Exit(0)
	}

	//ignoring rest of the args
	statusCodeToExit, err := strconv.ParseInt(strings.TrimSpace(args[0]), 10, 64)
	if err != nil {
		return CommandOutput{1, "", fmt.Errorf("string argument passed with exit")}
	}

	os.Exit(int(statusCodeToExit))

	// complier forced return
	return CommandOutput{0, "", nil}
}
