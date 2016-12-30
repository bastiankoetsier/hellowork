package cmd

import (
	"github.com/italolelis/hanu"
	log "github.com/Sirupsen/logrus"
)

var commandList []hanu.CommandInterface

type Command interface {
	Name() string
	Description() string
	Commands() []string
	Handler(conv hanu.ConversationInterface)
}

// Register adds a new command to commandList
func Register(command Command) {
	log.Debugf("% command registered", command.Name())
	cmds := command.Commands()
	for _, route := range cmds {
		commandList = append(commandList, hanu.NewCommand(command.Name(), command.Description(), route, command.Handler))
	}
}

// List returns commandList
func List() []hanu.CommandInterface {
	return commandList
}
