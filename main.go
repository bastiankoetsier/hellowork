package main

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/italolelis/hanu"
	"github.com/italolelis/hellowork/config"
	"github.com/italolelis/hellowork/cmd"
	"github.com/italolelis/hellowork/repo"
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

	bot, err := hanu.New(globalConfig.SlackToken)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Register(cmd.NewHi())
	cmd.Register(cmd.NewWhereIs(inMemoryRepo))

	cmdList := cmd.List()
	for _, command := range cmdList {
		bot.Register(command)
	}

	bot.Listen()

	//bot.Hear("(?i)I'm on (.*)").MessageHandler(ReasonHandler)
	//bot.Hear("(?i)I'm (.*)").MessageHandler(ReasonHandler)
	//bot.Run()
}

//func WhereIsEverybodyHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
//	users := service.WhereIsEverybody(time.Now())
//
//	msg := "These are the people who are out: \n"
//	for _, user := range users {
//		msg += user.String() + "\n"
//	}
//
//	bot.Reply(evt, msg, slackbot.WithTyping)
//}
//
//func ReasonHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
//	text := Message(evt.Text)
//	timable, err := text.WhatTimable()
//	if nil != err {
//		bot.Reply(evt, "I'm sorry I can't understand you", slackbot.WithTyping)
//	}
//
//	from := timable.From
//	to := timable.To
//
//	slackUser, err := bot.RTM.GetUserInfo(string(evt.User))
//	if nil != err {
//		log.Panic(err)
//	}
//
//	service.CreateStatus(slackUser, NewStatus("", from, to, Vacation))
//	if timable.HasOnlyFrom() {
//		bot.Reply(evt, "Ok and when will you be back?", slackbot.WithTyping)
//	} else {
//		bot.Reply(evt, fmt.Sprintf("Ok you are on vacations from %s until %s. Enjoy!", from.Format("02/01/2006"), to.Format("02/01/2006")), slackbot.WithTyping)
//	}
//}
