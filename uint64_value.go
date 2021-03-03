package strutlog

type uint64Value struct {
	value uint64
}

func (b uint64Value) Type() PrimitiveType {
	return UInt64Type
}

func (b uint64Value) Bool() bool {
	return defaultBoolVal
}

func (b uint64Value) String() string {
	return defaultStringVal
}

func (b uint64Value) Int() int {
	return defaultIntVal
}

func (b uint64Value) Int8() int8 {
	return defaultIntVal
}

func (b uint64Value) Int16() int16 {
	return defaultIntVal
}

func (b uint64Value) Int32() int32 {
	return defaultIntVal
}

func (b uint64Value) Int64() int64 {
	return defaultIntVal
}

func (b uint64Value) UInt() uint {
	return defaultIntVal
}

func (b uint64Value) UInt8() uint8 {
	return defaultIntVal
}

func (b uint64Value) UInt16() uint16 {
	return defaultIntVal
}

func (b uint64Value) UInt32() uint32 {
	return defaultIntVal
}

func (b uint64Value) UInt64() uint64 {
	return b.value
}

func (b uint64Value) Byte() byte {
	return defaultIntVal
}

func (b uint64Value) Rune() rune {
	return defaultIntVal
}

func (b uint64Value) Float32() float32 {
	return defaultFloatVal
}

func (b uint64Value) Float64() float64 {
	return defaultFloatVal
}

func (b uint64Value) Complex64() complex64 {
	return defaultComplex
}

func (b uint64Value) Complex128() complex128 {
	return defaultComplex
}

func (b uint64Value) Array() []Value {
	return []Value{}
}

func (b uint64Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b uint64Value) Strut() Fieldable {
	return defaultFieldable{}
}

func UInt64(value uint64) Value {
	return uint64Value{value}
}
