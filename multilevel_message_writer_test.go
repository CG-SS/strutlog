package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type countMessageWriter struct {
	count int
}

func (c *countMessageWriter) WriteMessage(_ Message) {
	c.count = c.count + 1
}

func TestMultiLevelMessageWriterWritesMessage(t *testing.T) {
	traceWriter := &countMessageWriter{}
	debugWriter := &countMessageWriter{}
	infoWriter := &countMessageWriter{}
	warnWriter := &countMessageWriter{}
	errorWriter := &countMessageWriter{}
	panicWriter := &countMessageWriter{}

	messageWriter := MultiLevelMessageWriter{
		TraceWriter: traceWriter,
		DebugWriter: debugWriter,
		InfoWriter:  infoWriter,
		WarnWriter:  warnWriter,
		ErrorWriter: errorWriter,
		PanicWriter: panicWriter,
	}

	messageWriter.WriteMessage(defaultMessage{
		level: TraceLevel,
	})
	assert.Equal(t, 1, traceWriter.count)

	messageWriter.WriteMessage(defaultMessage{
		level: DebugLevel,
	})
	assert.Equal(t, 1, debugWriter.count)

	messageWriter.WriteMessage(defaultMessage{
		level: InfoLevel,
	})
	assert.Equal(t, 1, infoWriter.count)

	messageWriter.WriteMessage(defaultMessage{
		level: WarnLevel,
	})
	assert.Equal(t, 1, warnWriter.count)

	messageWriter.WriteMessage(defaultMessage{
		level: ErrorLevel,
	})
	assert.Equal(t, 1, errorWriter.count)

	messageWriter.WriteMessage(defaultMessage{
		level: PanicLevel,
	})
	assert.Equal(t, 1, panicWriter.count)
}
