package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}

	_, err = validateCommand(input)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func validateCommand(input string) (Command, error) {
	tokens := strings.Split(input, SPACE)

	if len(tokens) == 0 {
		return Command{""}, nil // no command found
	}

	command := strings.TrimSpace(tokens[0])

	for _, c := range commands {
		if command == c.invocation {
			return c, nil
		}
	}

	return Command{""}, fmt.Errorf("%s: command not found", command)
}
