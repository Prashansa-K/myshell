package main

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
)

type commandFunction func([]string) (int, string, error) // int to describe success of failure, string to describe command output

type Command struct {
	invocation  string
	args        []string
	handlerFunc commandFunction
}

var validCommands = []Command{
	{"exit", nil, commands.Exit},
}
