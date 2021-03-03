package strutlog

type complex128Value struct {
	value complex128
}

func (b complex128Value) Type() PrimitiveType {
	return Complex128Type
}

func (b complex128Value) Bool() bool {
	return defaultBoolVal
}

func (b complex128Value) String() string {
	return defaultStringVal
}

func (b complex128Value) Int() int {
	return defaultIntVal
}

func (b complex128Value) Int8() int8 {
	return defaultIntVal
}

func (b complex128Value) Int16() int16 {
	return defaultIntVal
}

func (b complex128Value) Int32() int32 {
	return defaultIntVal
}

func (b complex128Value) Int64() int64 {
	return defaultIntVal
}

func (b complex128Value) UInt() uint {
	return defaultIntVal
}

func (b complex128Value) UInt8() uint8 {
	return defaultIntVal
}

func (b complex128Value) UInt16() uint16 {
	return defaultIntVal
}

func (b complex128Value) UInt32() uint32 {
	return defaultIntVal
}

func (b complex128Value) UInt64() uint64 {
	return defaultIntVal
}

func (b complex128Value) Byte() byte {
	return defaultIntVal
}

func (b complex128Value) Rune() rune {
	return defaultIntVal
}

func (b complex128Value) Float32() float32 {
	return defaultFloatVal
}

func (b complex128Value) Float64() float64 {
	return defaultFloatVal
}

func (b complex128Value) Complex64() complex64 {
	return defaultComplex
}

func (b complex128Value) Complex128() complex128 {
	return b.value
}

func (b complex128Value) Array() []Value {
	return []Value{}
}

func (b complex128Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b complex128Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Complex128(value complex128) Value {
	return complex128Value{value}
}
