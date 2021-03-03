package strutlog

type complex64Value struct {
	value complex64
}

func (b complex64Value) Type() PrimitiveType {
	return Complex64Type
}

func (b complex64Value) Bool() bool {
	return defaultBoolVal
}

func (b complex64Value) String() string {
	return defaultStringVal
}

func (b complex64Value) Int() int {
	return defaultIntVal
}

func (b complex64Value) Int8() int8 {
	return defaultIntVal
}

func (b complex64Value) Int16() int16 {
	return defaultIntVal
}

func (b complex64Value) Int32() int32 {
	return defaultIntVal
}

func (b complex64Value) Int64() int64 {
	return defaultIntVal
}

func (b complex64Value) UInt() uint {
	return defaultIntVal
}

func (b complex64Value) UInt8() uint8 {
	return defaultIntVal
}

func (b complex64Value) UInt16() uint16 {
	return defaultIntVal
}

func (b complex64Value) UInt32() uint32 {
	return defaultIntVal
}

func (b complex64Value) UInt64() uint64 {
	return defaultIntVal
}

func (b complex64Value) Byte() byte {
	return defaultIntVal
}

func (b complex64Value) Rune() rune {
	return defaultIntVal
}

func (b complex64Value) Float32() float32 {
	return defaultFloatVal
}

func (b complex64Value) Float64() float64 {
	return defaultFloatVal
}

func (b complex64Value) Complex64() complex64 {
	return b.value
}

func (b complex64Value) Complex128() complex128 {
	return defaultComplex
}

func (b complex64Value) Array() []Value {
	return []Value{}
}

func (b complex64Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b complex64Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Complex64(value complex64) Value {
	return complex64Value{value}
}
