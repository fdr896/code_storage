package core

import (
	"time"

	"github.com/rs/xid"
)

// Code - structer which contains information about particular code snippet
type Code struct {
	ID       string    `json:"id,omitempty"`
	Source   string    `json:"source,omitempty"`
	Language string    `json:"language,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
}

func NewCode(code *Code) Code {
	code.ID = xid.New().String()
	code.Date = time.Now()

	return *code
}
