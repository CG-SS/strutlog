package strutlog

import (
	"context"
	"os"
	"time"
)

// HookFn is a function that is evaluated right before the Message gets logged. You can inject fields and create your
// own Message implementation. If the function returns nil, then the message is later discarded.
type HookFn func(Message) Message

// defaultTimestampFieldName is used as the field name for the timestamp if FieldName is not set on TimestampConfig.
const defaultTimestampFieldName = "timestamp"

// defaultLevelFieldName is used as the field name for the level if FieldName is not set on LevelConfig.
const defaultLevelFieldName = "level"

// TimestampConfig contains all the configuration pertaining the timestamp on a Message.
type TimestampConfig struct {
	Type         TimestampType // Type of the timestamp.
	FieldName    string        // FieldName for the timestamp value.
	UTCEnabled   bool          // UTCEnabled makes the timestamp field use the UTC standard.
	CustomFormat string        // CustomFormat should be set if Type is CustomTimestamp.
}

// LevelConfig contains all the configuration pertaining the Message level field.
type LevelConfig struct {
	DisabledPrint bool         // DisabledPrint disables adding the level field if true.
	FieldName     string       // FieldName for the MessageLevel value.
	Value         MessageLevel // Value is the minimum MessageLevel the defaultLogger will print.
}

// DefaultLoggerConfig contains all the configuration pertaining a defaultLogger.
type DefaultLoggerConfig struct {
	BufferSize uint            // BufferSize of the internal Message channel. If set to 0, a non buffered channel will be used.
	SampleTime time.Duration   // SampleTime of the interval between writes. Minimum is a Nanosecond.
	Writer     MessageWriter   // Writer for the Messages.
	HookFns    []HookFn        // HookFns are the HookFn that will get run prior to the Message getting logged.
	Timestamp  TimestampConfig // Timestamp is the TimestampConfig used for this DefaultLoggerConfig.
	Level      LevelConfig     // Level is the LevelConfig used for this DefaultLoggerConfig.
}

// defaultLogger represents a standard logger that will receive message on the messageCh and log them accordingly.
// Most of the logic is actually handed by the writeHandler associated with this defaultLogger.
type defaultLogger struct {
	messageCh chan Message
}

// Trace creates a Message of MessageLevel TraceLevel.
func (d defaultLogger) Trace() Message {
	return defaultMessage{
		messageCh: d.messageCh,
		level:     TraceLevel,
	}
}

// Debug creates a Message of MessageLevel DebugLevel.
func (d defaultLogger) Debug() Message {
	return defaultMessage{
		messageCh: d.messageCh,
		level:     DebugLevel,
	}
}

// Info creates a Message of MessageLevel InfoLevel.
func (d defaultLogger) Info() Message {
	return defaultMessage{
		messageCh: d.messageCh,
		level:     InfoLevel,
	}
}

// Warn creates a Message of MessageLevel WarnLevel.
func (d defaultLogger) Warn() Message {
	return defaultMessage{
		messageCh: d.messageCh,
		level:     WarnLevel,
	}
}

// Error creates a Message of MessageLevel ErrorLevel.
func (d defaultLogger) Error() Message {
	return defaultMessage{
		messageCh: d.messageCh,
		level:     ErrorLevel,
	}
}

// Fatal creates a Message of MessageLevel FatalLevel.
func (d defaultLogger) Fatal() Message {
	return defaultMessage{
		messageCh: d.messageCh,
		level:     FatalLevel,
	}
}

// Panic creates a Message of MessageLevel PanicLevel.
func (d defaultLogger) Panic() Message {
	return defaultMessage{
		messageCh: d.messageCh,
		level:     PanicLevel,
	}
}

// NewDefaultLogger returns a default Logger that prints to os.Stdout as soon as possible and marshals Message using a
// JsonMarshaller, that, with JSON format.
// The internal writer will be stopped when the context.Context ctx is done.
func NewDefaultLogger(ctx context.Context) Logger {
	return NewLogger(ctx, DefaultLoggerConfig{
		BufferSize: 1024,
		Writer: SingleMessageWriter{
			Writer:  os.Stdout,
			Marshal: JsonMarshaller(-1, NoExponentFloatFormat),
		},
	})
}

// NewLogger returns a Logger with the given DefaultLoggerConfig cfg.
// The internal writer will be stopped when the context.Context ctx is done.
func NewLogger(ctx context.Context, cfg DefaultLoggerConfig) Logger {
	var messageCh chan Message
	if cfg.BufferSize == 0 {
		messageCh = make(chan Message)
	} else {
		messageCh = make(chan Message, cfg.BufferSize)
	}

	if cfg.Writer == nil {
		cfg.Writer = SingleMessageWriter{
			Writer:  os.Stdout,
			Marshal: JsonMarshaller(-1, NoExponentFloatFormat),
		}
	}

	if len(cfg.Timestamp.FieldName) == 0 {
		cfg.Timestamp.FieldName = defaultTimestampFieldName
	}
	if len(cfg.Level.FieldName) == 0 {
		cfg.Level.FieldName = defaultLevelFieldName
	}

	hookFns := cfg.HookFns
	if cfg.Timestamp.Type != DisabledTimestamp {
		hookFns = append(hookFns, applyTimestamp(cfg.Timestamp))
	}
	if !cfg.Level.DisabledPrint {
		hookFns = append(hookFns, applyLevel(cfg.Level))
	}

	go writeHandler{
		ctx:        ctx,
		writer:     cfg.Writer,
		messageCh:  messageCh,
		level:      cfg.Level.Value,
		hookFns:    hookFns,
		sampleTime: cfg.SampleTime,
	}.writeMsg()

	return defaultLogger{
		messageCh: messageCh,
	}
}

// writeHandler handles writing a Message to a MessageWriter.
// Effectively, most of the defaultLogger logic resides in here.
type writeHandler struct {
	ctx        context.Context
	writer     MessageWriter
	messageCh  <-chan Message
	level      MessageLevel
	hookFns    []HookFn
	sampleTime time.Duration
}

// writeMsg to the given MessageWriter.
// When sampleTime passes, we poll the messageCh. If a Message is present, we then check if the given Message has a
// greater MessageLevel than level, if it does, we proceed, otherwise that Message is discarded. Then, we run all the
// hookFns on the Message, if the Message is not nil after that, we write the message using writer.
func (w writeHandler) writeMsg() {
	sampleTimeNorm := w.sampleTime
	if sampleTimeNorm <= 0 {
		sampleTimeNorm = time.Nanosecond
	}
	ticker := time.NewTicker(sampleTimeNorm)
	defer ticker.Stop()

out:
	for {
		select {
		case <-w.ctx.Done():
			break out
		case <-ticker.C:
			for msg := range w.messageCh {
				if msg == nil || msg.Level() < w.level {
					continue
				}

				finalMsg := msg
				for _, fn := range w.hookFns {
					finalMsg = fn(finalMsg)
				}
				if finalMsg == nil {
					continue
				}

				w.writer.WriteMessage(finalMsg)
			}
		}
	}
}

// applyTimestamp generates a HookFn that will apply a timestamp field based on the TimestampConfig cfg.
func applyTimestamp(cfg TimestampConfig) HookFn {
	return func(message Message) Message {
		if message == nil {
			return nil
		}

		currentTime := time.Now()

		if cfg.UTCEnabled {
			currentTime = currentTime.UTC()
		}

		switch cfg.Type {
		case LayoutTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.Layout))
		case ANSICTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.ANSIC))
		case UnixDateTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.UnixDate))
		case RubyDateTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RubyDate))
		case RFC822Timestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RFC822))
		case RFC822ZTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RFC822Z))
		case RFC850Timestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RFC850))
		case RFC1123Timestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RFC1123))
		case RFC1123ZTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RFC1123Z))
		case RFC3339Timestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RFC3339))
		case RFC3339NanoTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.RFC3339Nano))
		case KitchenTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.Kitchen))
		case StampTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.Stamp))
		case StampMilliTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.StampMilli))
		case StampMicroTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.StampMicro))
		case StampNanoTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.StampNano))
		case DateTimeTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.DateTime))
		case DateOnlyTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.DateOnly))
		case TimeOnlyTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(time.TimeOnly))
		case UnixTimestamp:
			return message.Int64(cfg.FieldName, currentTime.Unix())
		case UnixNanoTimestamp:
			return message.Int64(cfg.FieldName, currentTime.UnixNano())
		case UnixMicroTimestamp:
			return message.Int64(cfg.FieldName, currentTime.UnixMicro())
		case UnixMilliTimestamp:
			return message.Int64(cfg.FieldName, currentTime.UnixMilli())
		case CustomTimestamp:
			return message.String(cfg.FieldName, currentTime.Format(cfg.CustomFormat))
		default:
			return message
		}
	}
}

// applyLevel returns a HookFn that will append a MessageLevel to the Message based on the given LevelConfig cfg.
func applyLevel(cfg LevelConfig) HookFn {
	return func(message Message) Message {
		if message == nil {
			return nil
		}

		return message.String(cfg.FieldName, message.Level().String())
	}
}
