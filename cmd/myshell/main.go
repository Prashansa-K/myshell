package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
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
			executeCommand(cmd)
		}
	}
}

func validateCommand(input string) (commands.Command, error) {
	tokens := strings.Split(input, SPACE)

	if len(tokens) == 0 {
		return commands.Command{}, nil // no command found
	}

	command := strings.TrimSpace(tokens[0])

	for _, c := range commands.ValidCommands {
		if command == c.Invocation {
			c.Args = tokens[1:]
			return c, nil
		}
	}

	return commands.Command{}, fmt.Errorf("%s: command not found", command)
}

func executeCommand(cmd commands.Command) {
	cmdOutput := cmd.HandlerFunc(cmd.Args)

	if cmdOutput.StatusCode != 0 {
		fmt.Println(cmdOutput.Error.Error())
		return
	}

	fmt.Print(cmdOutput.Output)
}
