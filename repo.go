package main

import (
	"time"
)

type Repository interface {
	Find(id string) *User
	FindAll() []*User
	FindAllOut(date time.Time) []*User
	FindAllByID(usernames []string, date time.Time) []*User
	Add(user *User)
	Remove(username UserID)
}

type InMemoryRepository struct {
	users map[UserID]*User
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{make(map[UserID]*User)}
}

func (r *InMemoryRepository) Find(id string) *User {
	user, exists := r.users[UserID(id)]

	if !exists {
		return nil
	}

	return user
}

func (r *InMemoryRepository) FindAll() []*User {
	var users []*User

	for _, user := range r.users {
		users = append(users, user)
	}

	return users
}

func (r *InMemoryRepository) FindAllOut(date time.Time) []*User {
	var users []*User

	for _, user := range r.users {
		if user.isAvailable(date) {
			users = append(users, user)
		}
	}

	return users
}

func (r *InMemoryRepository) FindAllByID(ids []string, date time.Time) []*User {
	var users []*User

	for _, id := range ids {
		user, exists := r.users[UserID(id)]
		if exists {
			users = append(users, user)
		}
	}

	return users
}

func (r *InMemoryRepository) Add(user *User) {
	r.users[user.ID] = user
}

func (r *InMemoryRepository) Remove(id UserID) {
	delete(r.users, id)
}
