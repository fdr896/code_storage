package core

import "time"

// Code - structer which contains information about particular code snippet
type Code struct {
	ID       string
	Source   string
	Language string
	Date     time.Time
	Tags     []string
}
