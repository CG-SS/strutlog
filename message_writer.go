package strutlog

// MessageMarshaller represents a function that can convert a Message into []bytes.
type MessageMarshaller func(Message) []byte

// MessageWriter represents a writer that knows how to write a Message.
type MessageWriter interface {
	WriteMessage(message Message)
}
