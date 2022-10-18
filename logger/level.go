package logger

type Level uint32

const (
	TraceLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (l *Level) String() string {
	switch *l {
	case TraceLevel:
		return "trace"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}

	return "unknown"
}

func (l *Level) Color() uint {
	switch *l {
	case TraceLevel:
		return gray
	case InfoLevel:
		return blue
	case WarnLevel:
		return yellow
	case ErrorLevel, FatalLevel:
		return red
	}

	return blue
}
