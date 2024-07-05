package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Exit(input string) (int, string, error) {
	tokens := strings.Split(input, " ")

	if len(tokens) == 1 { // only exit passed, assuming it as exit 0
		os.Exit(0)
	}

	statusCodeToExit, err := strconv.ParseInt(strings.TrimSpace(tokens[1]), 10, 64)
	if err != nil {
		return 1, "", fmt.Errorf("string argument passed with exit")
	}

	os.Exit(int(statusCodeToExit))

	// complier forced return
	return 0, "", nil
}
