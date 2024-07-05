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
	{"type", nil, Type},
}

// maintaining one more list to avoid init cycle
// if we use above struct, ValidCommand.type refers to Type func, which in turn refers to ValidCommands, thus creating a cycle
var validCommands = []string{"exit", "echo", "type"}
