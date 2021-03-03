package strutlog

type uint8Value struct {
	value uint8
}

func (b uint8Value) Type() PrimitiveType {
	return UInt8Type
}

func (b uint8Value) Bool() bool {
	return defaultBoolVal
}

func (b uint8Value) String() string {
	return defaultStringVal
}

func (b uint8Value) Int() int {
	return defaultIntVal
}

func (b uint8Value) Int8() int8 {
	return defaultIntVal
}

func (b uint8Value) Int16() int16 {
	return defaultIntVal
}

func (b uint8Value) Int32() int32 {
	return defaultIntVal
}

func (b uint8Value) Int64() int64 {
	return defaultIntVal
}

func (b uint8Value) UInt() uint {
	return defaultIntVal
}

func (b uint8Value) UInt8() uint8 {
	return b.value
}

func (b uint8Value) UInt16() uint16 {
	return defaultIntVal
}

func (b uint8Value) UInt32() uint32 {
	return defaultIntVal
}

func (b uint8Value) UInt64() uint64 {
	return defaultIntVal
}

func (b uint8Value) Byte() byte {
	return defaultIntVal
}

func (b uint8Value) Rune() rune {
	return defaultIntVal
}

func (b uint8Value) Float32() float32 {
	return defaultFloatVal
}

func (b uint8Value) Float64() float64 {
	return defaultFloatVal
}

func (b uint8Value) Complex64() complex64 {
	return defaultComplex
}

func (b uint8Value) Complex128() complex128 {
	return defaultComplex
}

func (b uint8Value) Array() []Value {
	return []Value{}
}

func (b uint8Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b uint8Value) Strut() Fieldable {
	return defaultFieldable{}
}

func UInt8(value uint8) Value {
	return uint8Value{value}
}
