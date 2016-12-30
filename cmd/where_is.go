package cmd

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/italolelis/hanu"
	"github.com/italolelis/hellowork/repo"
)

type UserParam struct {
	Param string
	Pattern *regexp.Regexp
}

func NewUserParam(conv hanu.ConversationInterface) (*UserParam, error) {
	param, err := conv.String("user")
	if nil != err || len(param) <= 0 {
		return nil, err
	}

	return &UserParam{param, regexp.MustCompile(`<@([a-zA-z0-9]+)>`)}, nil
}

func (u *UserParam) GetUserID() string {
	if u.Pattern.MatchString(u.Param) {
		results := u.Pattern.FindStringSubmatch(u.Param)
		return results[1]
	}

	return u.Param
}

func (u *UserParam) isEverybody() bool {
	return strings.ToLower(u.Param) == "everybody"
}

type WhereIs struct {
	repo repo.Repository
}

func NewWhereIs(repo repo.Repository) *WhereIs {
	return &WhereIs{repo}
}

func (c *WhereIs) Command() string {
	return "(?i)where is <user>"
}

func (c *WhereIs) Description() string {
	return "Finds if an user is available"
}

func (c *WhereIs) Handler(conv hanu.ConversationInterface) {
	userParam, err := NewUserParam(conv)
	if nil != err {
		log.Error(err)
		conv.Reply("I'm sorry I couldn't understand you")
	}

	if userParam.isEverybody() {
		var msg string
		users := c.repo.FindAllOut(time.Now())
		if len(users) > 0 {
			msg = "This are the people out: \n"
			for _, user := range users {
				msg += user.String()
			}
			conv.Reply(msg)
		} else {
			conv.Reply(fmt.Sprintf("As far as I know %s is available", userParam.Param))
		}
	} else {
		user := c.repo.Find(userParam.GetUserID())
		if nil == user {
			conv.Reply(fmt.Sprintf("As far as I know %s is available", userParam.Param))
		} else {
			conv.Reply(user.String())
		}
	}
}
