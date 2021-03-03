package strutlog

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type recorderMessageWriter struct {
	messages []Message
}

func (r *recorderMessageWriter) WriteMessage(message Message) {
	r.messages = append(r.messages, message)
}

func TestNewLoggerReturnsValidLoggerIfEmptyCfg(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	logger := NewLogger(ctx, DefaultLoggerConfig{})
	assert.NotNil(t, logger)
}

func TestDefaultLoggerLogs(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	writer := recorderMessageWriter{}
	waitFor := 10 * time.Nanosecond

	logger := NewLogger(ctx, DefaultLoggerConfig{
		SampleTime: waitFor,
		Writer:     &writer,
	})

	logger.Trace().Bool("test", true).String("name", "test").Send()
	logger.Debug().Bool("test", true).String("name", "test").Send()
	logger.Info().Bool("test", true).String("name", "test").Send()
	logger.Warn().Bool("test", true).String("name", "test").Send()
	logger.Error().Error("error", errors.New("oops")).Send()

	<-time.After(waitFor * 2)

	messages := writer.messages
	assert.Len(t, messages, 5)
	assert.Equal(t, messages[0].Level(), TraceLevel)
	assert.Equal(t, messages[1].Level(), DebugLevel)
	assert.Equal(t, messages[2].Level(), InfoLevel)
	assert.Equal(t, messages[3].Level(), WarnLevel)
	assert.Equal(t, messages[4].Level(), ErrorLevel)
}

func TestDefaultLoggerHookFnsFilterIfNil(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	writer := recorderMessageWriter{}
	waitFor := 10 * time.Nanosecond

	logger := NewLogger(ctx, DefaultLoggerConfig{
		SampleTime: waitFor,
		Writer:     &writer,
		HookFns: []HookFn{
			func(message Message) Message {
				return nil
			},
		},
	})

	logger.Trace().Bool("test", true).String("name", "test").Send()
	<-time.After(waitFor * 2)

	messages := writer.messages
	assert.Len(t, messages, 0)
}

func TestDefaultLoggerHookFnsInjectFields(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	writer := recorderMessageWriter{}
	waitFor := 10 * time.Nanosecond

	logger := NewLogger(ctx, DefaultLoggerConfig{
		SampleTime: waitFor,
		Writer:     &writer,
		HookFns: []HookFn{
			func(message Message) Message {
				return message.Bool("test", true)
			},
			func(message Message) Message {
				return message.String("country", "test")
			},
		},
		Timestamp: TimestampConfig{
			Type: DisabledTimestamp,
		},
		Level: LevelConfig{
			DisabledPrint: true,
		},
	})

	logger.Info().Int("i", 10).Send()
	<-time.After(waitFor * 2)

	messages := writer.messages
	assert.Len(t, messages, 1)

	injectedMessage := messages[0]
	fields := injectedMessage.Fields()
	assert.Len(t, fields, 3)
	assert.Equal(t, fields[0].Value.Int(), 10)
	assert.Equal(t, fields[1].Value.Bool(), true)
	assert.Equal(t, fields[2].Value.String(), "test")
}

type jsonTest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func TestLoggerLogsJsonMessage(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	waitFor := 10 * time.Nanosecond
	var buf bytes.Buffer

	logger := NewLogger(ctx, DefaultLoggerConfig{
		SampleTime: waitFor,
		Writer: SingleMessageWriter{
			Writer:  &buf,
			Marshal: JsonMarshaller(-1, NoExponentFloatFormat),
		},
		Timestamp: TimestampConfig{
			Type: DisabledTimestamp,
		},
		Level: LevelConfig{
			DisabledPrint: true,
		},
	})

	logger.Info().Int("x", 10).Int("y", 20).Send()
	<-time.After(waitFor * 2)

	var jsonVal jsonTest
	err := json.Unmarshal(buf.Bytes(), &jsonVal)
	assert.NoError(t, err)

	assert.Equal(t, jsonVal.X, 10)
	assert.Equal(t, jsonVal.Y, 20)
}
