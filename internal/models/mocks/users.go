package mocks

import (
	"time"

	"snippetbox.namruslan.com/internal/models"
)

const AuthenticatedEmail = "alice@example.com"
const AuthenticatedPassword = "pa$$word"

type UserModel struct {
}

var mockUser = models.User{
	ID:             1,
	Name:           "Alice",
	Email:          AuthenticatedEmail,
	HashedPassword: nil,
	Created:        time.Now(),
}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == AuthenticatedEmail && password == AuthenticatedPassword {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) GetById(id int) (models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return models.User{}, models.ErrNoRecord
	}
}
