package main

import (
	"time"

	"github.com/nlopes/slack"
)

type AvailabilityService struct {
	repo Repository
}

func NewAvailabilityService(repo Repository) *AvailabilityService {
	return &AvailabilityService{repo}
}

func (s *AvailabilityService) Available(id string) bool {
	return s.repo.Find(id) == nil
}

func (s *AvailabilityService) WhereIsEverybody(date time.Time) []*User {
	return s.repo.FindAllOut(date)
}

func (s *AvailabilityService) WhereIs(id string, date time.Time) *User {
	return s.repo.Find(id)
}

func (s *AvailabilityService) WhereAre(ids []string, date time.Time) []*User {
	return s.repo.FindAllByID(ids, date)
}

func (s *AvailabilityService) CreateStatus(slackUser *slack.User, status *Status) error {
	var user *User
	user = s.repo.Find(slackUser.ID)
	if nil == user {
		user = NewUser(UserID(slackUser.ID))
		user.Username = slackUser.Name
		s.repo.Add(user)
	}

	user.AddStatus(status)
	return nil
}
