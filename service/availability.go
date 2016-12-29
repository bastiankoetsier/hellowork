package service

import (
	"time"

	"github.com/nlopes/slack"
	"github.com/italolelis/hellowork/repo"
	"github.com/italolelis/hellowork/model"
)

type Availability struct {
	repo repo.Repository
}

func NewAvailability(repo repo.Repository) *Availability {
	return &Availability{repo}
}

func (s *Availability) Available(id string) bool {
	return s.repo.Find(id) == nil
}

func (s *Availability) WhereIsEverybody(date time.Time) []*model.User {
	return s.repo.FindAllOut(date)
}

func (s *Availability) WhereIs(id string, date time.Time) *model.User {
	return s.repo.Find(id)
}

func (s *Availability) WhereAre(ids []string, date time.Time) []*model.User {
	return s.repo.FindAllByID(ids, date)
}

func (s *Availability) CreateStatus(slackUser *slack.User, status *model.Status) error {
	var user *model.User
	user = s.repo.Find(slackUser.ID)
	if nil == user {
		user = model.NewUser(model.UserID(slackUser.ID))
		user.Username = slackUser.Name
		s.repo.Add(user)
	}

	user.AddStatus(status)
	return nil
}
