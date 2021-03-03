package strutlog

// Value represents a value in the log. The value include all Golang's primitive types, an array of Value, a
// map[Value]Value or a Fieldable.
// The idea is that, when a MessageMarshaller function is parsing a Value, it will know which PrimitiveType it is by
// using Type, and call the appropriate type function, for instance, if the PrimitiveType is StringType, the function
// would call String, in which String would return the proper internal value of Value.
// All the functions are meant to be safe, that is, if you wrongly call one of the value functions, they will return a
// default value, there is no 'panic' or error that is returned. For instance, if you call UInt on a Float32Type, it
// should return defaultIntVal.
type Value interface {
	Type() PrimitiveType
	Bool() bool
	String() string
	Int() int
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
	UInt() uint
	UInt8() uint8
	UInt16() uint16
	UInt32() uint32
	UInt64() uint64
	Byte() byte
	Rune() rune
	Float32() float32
	Float64() float64
	Complex64() complex64
	Complex128() complex128
	Array() []Value
	Map() map[Value]Value
	Strut() Fieldable
}
