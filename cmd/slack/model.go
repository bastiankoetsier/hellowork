package main

import (
	"fmt"
	"strings"
	"time"

	htime "github.com/italolelis/hellowork/time"
)

var (
	OutOfOffice Reason = "out of office"
	Remote      Reason = "working remote"
	Sick        Reason = "sick"
	Vacation    Reason = "vacation"
	WorkTrip    Reason = "work trip"
)

type UserID string
type Reason string

func ParseTime(date string) time.Time {
	now := time.Now()
	switch strings.ToLower(date) {
	case "today":
		return htime.Today()
	case "tomorrow":
		return htime.Tomorrow()
	case "yesterday":
		return htime.Yesterday()
	case "next week":
		return htime.AddWeek(now)
	case "next month":
	default:
		return now
	}

	return now
}

func ParseReason(reason string) Reason {
	switch strings.ToLower(reason) {
	case "out of office":
		return OutOfOffice
	case "remote":
		return Remote
	case "sick":
		return Sick
	case "vacation":
		return Vacation
	case "work trip":
		return WorkTrip
	default:
		return OutOfOffice
	}
}

type User struct {
	ID       UserID
	Username string
	Statuses []*Status
}

func NewUser(id UserID) *User {
	return &User{ID: id, Statuses: make([]*Status, 0)}
}

func (u *User) AddStatus(status *Status) {
	u.Statuses = append(u.Statuses, status)
}

func (u *User) GetStatus() *Status {
	return u.Statuses[len(u.Statuses) - 1]
}

func (u *User) isAvailable(date time.Time) bool {
	var available bool
	for _, status := range u.Statuses {
		available = !status.isValid(date)
	}

	return available
}

func (u *User) String() string {
	from := u.GetStatus().From
	to := u.GetStatus().To
	return fmt.Sprintf("<@%s> is out from %s until %s (%s)", u.ID, from.Format("02/01/2006"), to.Format("Monday"), to.Format("02/01/2006"))
}

type Status struct {
	Description string
	From        time.Time
	To          time.Time
	Reason      Reason
}

func NewStatus(description string, from time.Time, to time.Time, reason Reason) *Status {
	return &Status{description, from, to, reason}
}

func (s *Status) isValid(date time.Time) bool {
	before := s.From.Before(date)
	after := s.To.After(date)

	return !before && !after
}
