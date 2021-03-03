package strutlog

// MessageLevel of a given Message. It represents a message level in terms of awareness.
type MessageLevel int

const (
	TraceLevel MessageLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

// String representation of the given MessageLevel.
func (l MessageLevel) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	default:
		return defaultStringVal
	}
}
