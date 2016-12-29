package cmd

import "github.com/italolelis/hanu"

type Hi struct{}

func NewHi() *Hi {
	return &Hi{}
}

func (s *Hi) Command() string {
	return "hi"
}

func (s *Hi) Description() string {
	return "Greeting someone"
}

func (s *Hi) Handler(conv hanu.ConversationInterface) {
	conv.Reply("Oh hello!")
}
