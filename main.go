package main

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/italolelis/hanu"
	"github.com/italolelis/hellowork/cmd"
	"github.com/italolelis/hellowork/config"
	"github.com/italolelis/hellowork/repo"
	"github.com/nlopes/slack"
)

var (
	err          error
	globalConfig *config.Specification
)

// initializes the global configuration
func init() {
	globalConfig, err = config.LoadEnv()
	if nil != err {
		log.Panic(err.Error())
	}
}

// initializes the basic configuration for the log wrapper
func init() {
	level, err := log.ParseLevel(strings.ToLower(globalConfig.LogLevel))
	if err != nil {
		log.Error("Error getting level", err)
	}

	log.SetLevel(level)
}

func main() {
	inMemoryRepo := repo.NewInMemory()
	client := slack.New(globalConfig.SlackToken)
	bot, err := hanu.NewWithConnection(hanu.NewSlackRTMConnection(client))
	if err != nil {
		log.Fatal(err)
	}

	cmd.Register(cmd.NewHi())
	cmd.Register(cmd.NewWhereIs(inMemoryRepo))
	cmd.Register(cmd.NewStatus(client, inMemoryRepo))

	cmdList := cmd.List()
	for _, command := range cmdList {
		bot.Register(command)
	}

	bot.Listen()
}
