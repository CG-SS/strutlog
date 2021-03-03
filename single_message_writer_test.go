package strutlog

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestSingleMessageWriterWritesMessage(t *testing.T) {
	testStr := "testing string write"
	var bytesWriter bytes.Buffer

	msgWriter := SingleMessageWriter{
		Writer: &bytesWriter,
		Marshal: func(message Message) []byte {
			return []byte(testStr)
		},
		ErrorHandler: nil,
	}

	msgWriter.WriteMessage(defaultMessage{})

	assert.Equal(t, testStr, bytesWriter.String())
}

type errorWriter struct {
}

func (e errorWriter) Write(_ []byte) (n int, err error) {
	return 0, errors.New("error")
}

func TestSingleMessageWriterHandleError(t *testing.T) {
	atomicBool := atomic.Bool{}
	atomicBool.Store(false)

	msgWriter := SingleMessageWriter{
		Writer: errorWriter{},
		Marshal: func(message Message) []byte {
			return []byte("test")
		},
		ErrorHandler: func(_ Message, _ error) {
			atomicBool.Store(true)
		},
	}

	msgWriter.WriteMessage(defaultMessage{})

	assert.True(t, atomicBool.Load())
}
