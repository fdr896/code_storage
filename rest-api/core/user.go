package core

import (
	"errors"
	"strings"
)

// User stores information about particular user.
type User struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

var (
	ErrNoSuchUser        = errors.New("no user with required login")
	ErrEmptyLogin        = errors.New("user's login field is empty")
	ErrEmptyPassword     = errors.New("new password should contain at least one character")
	ErrUserWithSameLogin = errors.New("user with required login is already registered")
)

// Valid checks user login for being blank.
func (us *User) Valid() bool {
	us.Login = strings.TrimSpace(us.Login)

	return us.Login != ""
}

// UserStorage represents interface which storage for user's info should satisfies.
type UserStorage interface {
	AddUser(user *User) error
	HasPassword(login string) (bool, error)
	SetPassword(login string, newPassword string) error
	ComparePassword(login string, receivedPassoword string) (bool, error)
}
