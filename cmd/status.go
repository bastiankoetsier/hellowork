package cmd

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/italolelis/hanu"
	"github.com/italolelis/hellowork/model"
	"github.com/italolelis/hellowork/repo"
	"github.com/nlopes/slack"
)

type Status struct {
	client *slack.Client
	repo   repo.Repository
}

func NewStatus(client *slack.Client, repo repo.Repository) *Status {
	return &Status{client, repo}
}

func (s *Status) Command() string {
	return `(?i)I'm on <status>(.*?)`
}

func (s *Status) Description() string {
	return "Creates a status for you"
}

func (s *Status) Handler(conv hanu.ConversationInterface) {
	statusParam, err := conv.String("status")
	timableParam, err := conv.Match(1)
	timable, err := model.NewTimableMention(timableParam)
	if nil == timable || nil != err {
		conv.Reply("I'm sorry I can't understand you")
		return
	}

	from := timable.From
	to := timable.To

	slackUser, err := s.client.GetUserInfo(conv.Message().UserID)
	if nil != err {
		log.Panic(err)
	}


	s.createStatus(slackUser, model.NewStatus("", from, to, model.ParseReason(statusParam)))
	if timable.HasOnlyFrom() {
		conv.Reply("Ok and when will you be back?")
	} else {
		conv.Reply(fmt.Sprintf("Ok you are on vacations from %s until %s. Enjoy!", from.Format("02/01/2006"), to.Format("02/01/2006")))
	}
}

func (s *Status) createStatus(slackUser *slack.User, status *model.Status) {
	var user *model.User
	user = s.repo.Find(slackUser.ID)
	if nil == user {
		user = model.NewUser(model.UserID(slackUser.ID))
		user.Username = slackUser.Name
	}

	user.AddStatus(status)
	s.repo.Add(user)
}
