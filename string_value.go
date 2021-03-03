package strutlog

type stringValue struct {
	value string
}

func (b stringValue) Type() PrimitiveType {
	return StringType
}

func (b stringValue) Bool() bool {
	return defaultBoolVal
}

func (b stringValue) String() string {
	return b.value
}

func (b stringValue) Int() int {
	return defaultIntVal
}

func (b stringValue) Int8() int8 {
	return defaultIntVal
}

func (b stringValue) Int16() int16 {
	return defaultIntVal
}

func (b stringValue) Int32() int32 {
	return defaultIntVal
}

func (b stringValue) Int64() int64 {
	return defaultIntVal
}

func (b stringValue) UInt() uint {
	return defaultIntVal
}

func (b stringValue) UInt8() uint8 {
	return defaultIntVal
}

func (b stringValue) UInt16() uint16 {
	return defaultIntVal
}

func (b stringValue) UInt32() uint32 {
	return defaultIntVal
}

func (b stringValue) UInt64() uint64 {
	return defaultIntVal
}

func (b stringValue) Byte() byte {
	return defaultIntVal
}

func (b stringValue) Rune() rune {
	return defaultIntVal
}

func (b stringValue) Float32() float32 {
	return defaultFloatVal
}

func (b stringValue) Float64() float64 {
	return defaultFloatVal
}

func (b stringValue) Complex64() complex64 {
	return defaultComplex
}

func (b stringValue) Complex128() complex128 {
	return defaultComplex
}

func (b stringValue) Array() []Value {
	return []Value{}
}

func (b stringValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b stringValue) Strut() Fieldable {
	return defaultFieldable{}
}

func String(value string) Value {
	return stringValue{value}
}
