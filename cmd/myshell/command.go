package main

type Command struct {
	invocation string
}

var commands = []Command{
	Command{"ls"}, Command{"echo"},
}
