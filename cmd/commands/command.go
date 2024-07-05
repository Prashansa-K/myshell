package commands

type commandFunction func([]string) CommandOutput

type Command struct {
	Invocation  string
	Args        []string
	HandlerFunc commandFunction
}

type CommandOutput struct {
	StatusCode int
	Output     string
	Error      error
}

var ValidCommands = []Command{
	{"exit", nil, Exit},
	{"echo", nil, Echo},
}
