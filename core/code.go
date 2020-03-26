package core

import (
	"errors"
	"strings"
	"time"

	"github.com/rs/xid"
)

var (
	ErrNotFound      = errors.New("code was not found")
	ErrEmptySource   = errors.New("code's source is empty")
	ErrEmptyLanguage = errors.New("code's language is empty")
)

// Code contains information about particular code snippet.
type Code struct {
	ID       string    `json:"id,omitempty"`
	Source   string    `json:"source,omitempty"`
	Language string    `json:"language,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
}

// NewCode adds id and date to instance of Code and returns it.
func (code *Code) NewCode() error {
	code.Source = strings.TrimSpace(code.Source)
	if code.Source == "" {
		return ErrEmptySource
	}

	code.Language = strings.TrimSpace(code.Language)
	if code.Language == "" {
		return ErrEmptyLanguage
	}

	code.ID = xid.New().String()
	code.Date = time.Now()

	return nil
}

// CodeStorage represents storage interface.
type CodeStorage interface {
	Get(id string) (*Code, error)
	GetAll() ([]*Code, error)
	Add(code *Code) error
	Delete(id string) error
	Close()
}
