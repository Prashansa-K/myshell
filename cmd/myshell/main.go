package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	shellReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := shellReader.ReadString('\n')
		if err != nil {
			return
		}

		cmd, err := validateCommand(input)
		if err != nil {
			fmt.Println(err)
		} else {
			executeCommand(cmd, input)
		}
	}
}

func validateCommand(input string) (Command, error) {
	tokens := strings.Split(input, SPACE)

	if len(tokens) == 0 {
		return Command{}, nil // no command found
	}

	command := strings.TrimSpace(tokens[0])

	for _, c := range validCommands {
		if command == c.invocation {
			return c, nil
		}
	}

	return Command{}, fmt.Errorf("%s: command not found", command)
}

func executeCommand(cmd Command, input string) {
	// statusCode, output, error :=
	cmd.handlerFunc(input)
}
