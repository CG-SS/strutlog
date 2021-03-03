package strutlog

type uint32Value struct {
	value uint32
}

func (b uint32Value) Type() PrimitiveType {
	return UInt32Type
}

func (b uint32Value) Bool() bool {
	return defaultBoolVal
}

func (b uint32Value) String() string {
	return defaultStringVal
}

func (b uint32Value) Int() int {
	return defaultIntVal
}

func (b uint32Value) Int8() int8 {
	return defaultIntVal
}

func (b uint32Value) Int16() int16 {
	return defaultIntVal
}

func (b uint32Value) Int32() int32 {
	return defaultIntVal
}

func (b uint32Value) Int64() int64 {
	return defaultIntVal
}

func (b uint32Value) UInt() uint {
	return defaultIntVal
}

func (b uint32Value) UInt8() uint8 {
	return defaultIntVal
}

func (b uint32Value) UInt16() uint16 {
	return defaultIntVal
}

func (b uint32Value) UInt32() uint32 {
	return b.value
}

func (b uint32Value) UInt64() uint64 {
	return defaultIntVal
}

func (b uint32Value) Byte() byte {
	return defaultIntVal
}

func (b uint32Value) Rune() rune {
	return defaultIntVal
}

func (b uint32Value) Float32() float32 {
	return defaultFloatVal
}

func (b uint32Value) Float64() float64 {
	return defaultFloatVal
}

func (b uint32Value) Complex64() complex64 {
	return defaultComplex
}

func (b uint32Value) Complex128() complex128 {
	return defaultComplex
}

func (b uint32Value) Array() []Value {
	return []Value{}
}

func (b uint32Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b uint32Value) Strut() Fieldable {
	return defaultFieldable{}
}

func UInt32(value uint32) Value {
	return uint32Value{value}
}
