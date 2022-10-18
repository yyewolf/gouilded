package logger

type Logger struct {
	// Minimum level to display by the logger. By default it is set to info
	Level Level

	// The formatter for the logger. By default it is TextFormatter
	Formatter Formatter
}

func NewLogger() *Logger {
	return &Logger{Level: InfoLevel, Formatter: NewTextFormatter()}
}

func (l *Logger) Trace(message string) {
	e := NewEntry(TraceLevel, message, l)
	e.Display()
}

func (l *Logger) Info(message string) {
	e := NewEntry(InfoLevel, message, l)
	e.Display()
}

func (l *Logger) Warn(message string) {
	e := NewEntry(WarnLevel, message, l)
	e.Display()
}

func (l *Logger) Error(message string) {
	e := NewEntry(ErrorLevel, message, l)
	e.Display()
}

func (l *Logger) Fatal(message string) {
	e := NewEntry(FatalLevel, message, l)
	e.Display()
}

func (l *Logger) TraceString(message string) string {
	e := NewEntry(TraceLevel, message, l)
	return e.ToString()
}

func (l *Logger) InfoString(message string) string {
	e := NewEntry(InfoLevel, message, l)
	return e.ToString()
}

func (l *Logger) WarnString(message string) string {
	e := NewEntry(WarnLevel, message, l)
	return e.ToString()
}

func (l *Logger) ErrorString(message string) string {
	e := NewEntry(ErrorLevel, message, l)
	return e.ToString()
}

func (l *Logger) FatalString(message string) string {
	e := NewEntry(FatalLevel, message, l)
	return e.ToString()
}
