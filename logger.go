package strutlog

// Logger is basically a Message factory. The idea is that a Logger would act as the first element of the factory, and
// after the Message is sent for logging, it would then effectively log it.
type Logger interface {
	Trace() Message
	Debug() Message
	Info() Message
	Warn() Message
	Error() Message
	Fatal() Message
	Panic() Message
}
