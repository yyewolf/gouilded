package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type loggerTest struct {
	Logger *Logger
}

func TestLogger_New(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	assert.Equal(t, v.Logger.Level, InfoLevel)
}

func TestLogger_Trace(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(TraceLevel, "test", v.Logger)

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("\x1b[37m[TRACE]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_TraceNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(TraceLevel, "test", v.Logger)
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("[TRACE] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_Info(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(InfoLevel, "test", v.Logger)

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("\x1b[36m[INFO]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_InfoNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(InfoLevel, "test", v.Logger)
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("[INFO] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_Warn(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(WarnLevel, "test", v.Logger)

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("\x1b[33m[WARN]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_WarnNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(WarnLevel, "test", v.Logger)
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("[WARN] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_Error(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(ErrorLevel, "test", v.Logger)

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("\x1b[31m[ERROR]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_ErrorNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(ErrorLevel, "test", v.Logger)
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("[ERROR] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_Fatal(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(FatalLevel, "test", v.Logger)

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("\x1b[31m[FATAL]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_FatalNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(FatalLevel, "test", v.Logger)
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	s, err := v.Logger.Formatter.Format(e)

	assert.Equal(t, err, nil)
	assert.Equal(t, s, fmt.Sprintf("[FATAL] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_TraceStringDisabled(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	e := NewEntry(TraceLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, "")
}

func TestLogger_TraceString(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel

	e := NewEntry(TraceLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("\x1b[37m[TRACE]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_TraceStringNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	e := NewEntry(TraceLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("[TRACE] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_InfoStringDisabled(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = ErrorLevel

	e := NewEntry(InfoLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, "")
}

func TestLogger_InfoString(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel

	e := NewEntry(InfoLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("\x1b[36m[INFO]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_InfoStringNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	e := NewEntry(InfoLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("[INFO] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_WarnStringDisabled(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = ErrorLevel

	e := NewEntry(WarnLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, "")
}

func TestLogger_WarnString(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel

	e := NewEntry(WarnLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("\x1b[33m[WARN]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_WarnStringNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	e := NewEntry(WarnLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("[WARN] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_ErrorStringDisabled(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = FatalLevel

	e := NewEntry(ErrorLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, "")
}

func TestLogger_ErrorString(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel

	e := NewEntry(ErrorLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("\x1b[31m[ERROR]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_ErrorStringNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	e := NewEntry(ErrorLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("[ERROR] %s test", e.Time.Format("RFC3339")))
}

func TestLogger_FatalString(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel

	e := NewEntry(FatalLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("\x1b[31m[FATAL]\x1b[0m %s test", e.Time.Format("RFC3339")))
}

func TestLogger_FatalStringNoColor(t *testing.T) {
	v := loggerTest{
		Logger: NewLogger(),
	}

	v.Logger.Level = TraceLevel
	v.Logger.Formatter = TextFormatter{DisableColors: true}

	e := NewEntry(FatalLevel, "test", v.Logger)

	s := e.ToString()

	assert.Equal(t, s, fmt.Sprintf("[FATAL] %s test", e.Time.Format("RFC3339")))
}
