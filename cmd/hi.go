package cmd

import "github.com/italolelis/hanu"

type Hi struct{}

func NewHi() *Hi {
	return &Hi{}
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
	conv.Reply("Oh hello!")
}
