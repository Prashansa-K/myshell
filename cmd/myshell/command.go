package main

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
)

type commandFunction func(string) (int, string, error) // int to describe success of failure, string to describe command output

type Command struct {
	invocation  string
	handlerFunc commandFunction
}

var validCommands = []Command{
	{"ls", nil},
	{"echo", nil},
	{"exit", commands.Exit},
}
