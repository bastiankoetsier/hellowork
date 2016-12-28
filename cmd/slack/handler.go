package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/italolelis/hanu"
	"time"
	"regexp"
	"fmt"
)

func init() {
	Register("hi", "Greeting someone", HelloHandler)
	Register("where is <user>", "Finds if an user is available", WhereIsHandler)
}

func HelloHandler(conv hanu.ConversationInterface) {
	conv.Reply("Oh hello!")
}

func WhereIsHandler(conv hanu.ConversationInterface) {
	userID, err := conv.StripString("user", regexp.MustCompile(`<@([a-zA-z0-9]+)>`))
	if nil != err || len(userID) <= 0 {
		log.Error(err)
		conv.Reply("I'm sorry I couldn't understand you")
	}

	user := service.WhereIs(userID[1], time.Now())
	if nil == user {
		conv.Reply(fmt.Sprintf("As far as I know %s is available", userID[0]))
	} else {
		conv.Reply(user.String())
	}
}
