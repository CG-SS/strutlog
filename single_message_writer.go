package strutlog

import "io"

// ErrorHandlerFn represents a function that will get called in case of an error happens during writing a Message.
type ErrorHandlerFn func(Message, error)

// SingleMessageWriter is a MessageWriter that writer that writes a Message to the given Writer.
type SingleMessageWriter struct {
	Writer       io.Writer
	Marshal      MessageMarshaller
	ErrorHandler ErrorHandlerFn
}

// WriteMessage to the given Writer, and, if ErrorHandler is defined, calls it in case writing fails.
func (s SingleMessageWriter) WriteMessage(message Message) {
	bytes := s.Marshal(message)
	_, err := s.Writer.Write(bytes)

	if err != nil && s.ErrorHandler != nil {
		s.ErrorHandler(message, err)
	}
}
