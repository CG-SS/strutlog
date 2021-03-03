package strutlog

type uint16Value struct {
	value uint16
}

func (b uint16Value) Type() PrimitiveType {
	return UInt16Type
}

func (b uint16Value) Bool() bool {
	return defaultBoolVal
}

func (b uint16Value) String() string {
	return defaultStringVal
}

func (b uint16Value) Int() int {
	return defaultIntVal
}

func (b uint16Value) Int8() int8 {
	return defaultIntVal
}

func (b uint16Value) Int16() int16 {
	return defaultIntVal
}

func (b uint16Value) Int32() int32 {
	return defaultIntVal
}

func (b uint16Value) Int64() int64 {
	return defaultIntVal
}

func (b uint16Value) UInt() uint {
	return defaultIntVal
}

func (b uint16Value) UInt8() uint8 {
	return defaultIntVal
}

func (b uint16Value) UInt16() uint16 {
	return b.value
}

func (b uint16Value) UInt32() uint32 {
	return defaultIntVal
}

func (b uint16Value) UInt64() uint64 {
	return defaultIntVal
}

func (b uint16Value) Byte() byte {
	return defaultIntVal
}

func (b uint16Value) Rune() rune {
	return defaultIntVal
}

func (b uint16Value) Float32() float32 {
	return defaultFloatVal
}

func (b uint16Value) Float64() float64 {
	return defaultFloatVal
}

func (b uint16Value) Complex64() complex64 {
	return defaultComplex
}

func (b uint16Value) Complex128() complex128 {
	return defaultComplex
}

func (b uint16Value) Array() []Value {
	return []Value{}
}

func (b uint16Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b uint16Value) Strut() Fieldable {
	return defaultFieldable{}
}

func UInt16(value uint16) Value {
	return uint16Value{value}
}
