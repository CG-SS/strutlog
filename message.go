package strutlog

// Message represents a Logger message. The intention is that whoever implements the Message interface would use a
// factory pattern in order to build a Message.
type Message interface {
	Level() MessageLevel
	Fields() []Field
	Bool(key string, value bool) Message
	String(key string, value string) Message
	Int(key string, value int) Message
	Int8(key string, value int8) Message
	Int16(key string, value int16) Message
	Int32(key string, value int32) Message
	Int64(key string, value int64) Message
	UInt(key string, value uint) Message
	UInt8(key string, value uint8) Message
	UInt16(key string, value uint16) Message
	UInt32(key string, value uint32) Message
	UInt64(key string, value uint64) Message
	Byte(key string, value byte) Message
	Rune(key string, value rune) Message
	Float32(key string, value float32) Message
	Float64(key string, value float64) Message
	Complex64(key string, value complex64) Message
	Complex128(key string, value complex128) Message
	Array(key string, value []Value) Message
	Map(key string, value map[Value]Value) Message
	Strut(key string, value Fieldable) Message
	Error(key string, value error) Message
	Send() Message
}
