package cmd

import (
	"github.com/italolelis/hanu"
	"math/rand"
)

type Hi struct {
	greetings []string
}

func NewHi() *Hi {
	return &Hi{
		greetings: []string{"Hi", "Hello", "Hello, good to see you around", "Hey"},
	}
}

func (s *Hi) Commands() []string {
	return []string{
		"(hi|hello|hello there)",
	}
}

func (s *Hi) Name() string {
	return "Hi"
}

func (s *Hi) Description() string {
	return "Greeting someone"
}

func (s *Hi) Handler(conv hanu.ConversationInterface) {
	i := rand.Intn(len(s.greetings))
	conv.Reply(s.greetings[i])
}
