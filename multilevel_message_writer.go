package strutlog

// MultiLevelMessageWriter is a MessageWriter that writes a Message to a given MessageWriter based on their MessageLevel.
type MultiLevelMessageWriter struct {
	TraceWriter MessageWriter
	DebugWriter MessageWriter
	InfoWriter  MessageWriter
	WarnWriter  MessageWriter
	ErrorWriter MessageWriter
	PanicWriter MessageWriter
}

// WriteMessage a Message based on their MessageLevel.
// If it's InfoLevel, it will write using InfoWriter, and so on.
func (m MultiLevelMessageWriter) WriteMessage(message Message) {
	switch message.Level() {
	case TraceLevel:
		m.TraceWriter.WriteMessage(message)
	case DebugLevel:
		m.DebugWriter.WriteMessage(message)
	case InfoLevel:
		m.InfoWriter.WriteMessage(message)
	case WarnLevel:
		m.WarnWriter.WriteMessage(message)
	case ErrorLevel:
		m.ErrorWriter.WriteMessage(message)
	case PanicLevel:
		m.PanicWriter.WriteMessage(message)
	}
}
