package model

import (
	"regexp"
	"time"
	"errors"
)

var (
	timablePattern     = regexp.MustCompile(`(?i)(from|since)\s(today|yesterday|tomorrow|next week|next month)\s(until|till|to)\s(today|yesterday|tomorrow|next week|next month)`)
	fromTimablePattern = regexp.MustCompile(`(?i)(from|since)\s(today|yesterday|tomorrow|next week|next month)`)
	toTimablePattern   = regexp.MustCompile(`(?i)(until|till|to)\s(today|yesterday|tomorrow|next week|next month)`)
	defaultPattern     = regexp.MustCompile(`(?i)(today|yesterday|tomorrow|next week|next month)`)
)

var (
	ErrPatternDoestMatch = errors.New("couldn't match a pattern")
)

type TimableMention struct {
	From    time.Time
	HasFrom bool
	To      time.Time
	HasTo   bool
}

func NewTimableMention(msg string) (*TimableMention, error) {
	var results [][]string
	switch {
	case timablePattern.MatchString(msg):
		results = timablePattern.FindAllStringSubmatch(msg, -1)
		return &TimableMention{
			From:    ParseTime(results[0][0]),
			HasFrom: len(results[0][0]) > 0,
			To:      ParseTime(results[1][0]),
			HasTo:   len(results[1][0]) > 0,
		}, nil
	case fromTimablePattern.MatchString(msg):
		results = fromTimablePattern.FindAllStringSubmatch(msg, -1)
		return &TimableMention{
			From:    ParseTime(results[0][2]),
			HasFrom: true,
			HasTo:   false,
		}, nil
	case toTimablePattern.MatchString(msg):
		results = toTimablePattern.FindAllStringSubmatch(msg, -1)
		return &TimableMention{
			From:    time.Now(),
			HasFrom: true,
			To:      ParseTime(results[0][2]),
			HasTo:   true,
		}, nil
	case defaultPattern.MatchString(msg):
		results = defaultPattern.FindAllStringSubmatch(msg, -1)
		return &TimableMention{
			From:    ParseTime(results[0][0]),
			HasFrom: true,
			HasTo:   false,
		}, nil
	}

	return nil, ErrPatternDoestMatch
}

func (t *TimableMention) HasOnlyFrom() bool {
	return t.HasFrom && !t.HasTo
}
