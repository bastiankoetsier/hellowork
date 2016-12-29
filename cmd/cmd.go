package cmd

import "github.com/italolelis/hanu"

var commandList []hanu.CommandInterface

type Command interface {
	Command() string
	Description() string
	Handler(conv hanu.ConversationInterface)
}

// Register adds a new command to commandList
func Register(command Command) {
	commandList = append(commandList, hanu.NewCommand(command.Command(), command.Description(), command.Handler))
}

// List returns commandList
func List() []hanu.CommandInterface {
	return commandList
}
