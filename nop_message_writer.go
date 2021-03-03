package strutlog

type NopMessageWriter struct {
}

func (n NopMessageWriter) WriteMessage(_ Message) {}
