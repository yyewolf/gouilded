package logger

import (
	"fmt"
	"strings"
)

type Formatter interface {
	Format(*Entry) (string, error)
}

type TextFormatter struct {
	// Disable colors for the formatter. By default set to false
	DisableColors bool
}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{DisableColors: false}
}

func (f TextFormatter) Format(e *Entry) (string, error) {
	var levelText string
	var levelColor uint
	var format string

	levelText = strings.ToUpper(e.Level.String())

	if f.DisableColors {
		format = fmt.Sprintf("[%s] %s %s", levelText, e.Time.Format("RFC3339"), e.Message)
	} else {
		levelColor = e.Level.Color()
		format = fmt.Sprintf("\x1b[%dm[%s]\x1b[0m %s %s", levelColor, levelText, e.Time.Format("RFC3339"), e.Message)
	}

	return format, nil
}
