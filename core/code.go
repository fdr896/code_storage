package core

import (
	"errors"
	"net/http"
	"time"

	"github.com/rs/xid"
)

// Error represents a type with which errors in http requests handle
type Error struct {
	StatusCode   int
	ErrorMessage error
}

var (
	CodeDoesNotExist = Error{http.StatusNotFound, errors.New("Code with required id does not exist")}
	UnsupportedJSON  = Error{http.StatusNotAcceptable, errors.New("Received JSON doesn't not satisfies to requested conditions")}
)

// Code - structer which contains information about particular code snippet
type Code struct {
	ID       string    `json:"id,omitempty"`
	Source   string    `json:"source,omitempty"`
	Language string    `json:"language,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
}

// NewCode adds id and date to instance of Code and returns it
func NewCode(code *Code) Code {
	code.ID = xid.New().String()
	code.Date = time.Now()

	return *code
}

// CheckCode checks if Source and language filed of Code instance are not empty
func CheckCode(code *Code) bool {
	return code.Language != "" && code.Source != ""
}
