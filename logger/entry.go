package logger

import (
	"fmt"
	"time"
)

type Entry struct {
	// Parent logger of the entry
	Logger *Logger

	// Time when the entry appeared
	Time time.Time

	// Message of the entry
	Message string

	// Level of the entry
	Level Level
}

// Return a empty entry with the actual time.
func NewEntry(level Level, message string, logger *Logger) *Entry {
	return &Entry{Time: time.Now(), Level: level, Message: message, Logger: logger}
}

func (e *Entry) HasLevel() bool {
	return e.Level >= e.Logger.Level
}

func (e *Entry) Display() {
	if e.HasLevel() {
		s, err := e.Logger.Formatter.Format(e)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", s)
	}
}

func (e *Entry) ToString() string {
	if e.HasLevel() {
		s, err := e.Logger.Formatter.Format(e)

		if err != nil {
			panic(err)
		}

		return s
	}

	return ""
}
