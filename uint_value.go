package strutlog

type uintValue struct {
	value uint
}

func (b uintValue) Type() PrimitiveType {
	return UIntType
}

func (b uintValue) Bool() bool {
	return defaultBoolVal
}

func (b uintValue) String() string {
	return defaultStringVal
}

func (b uintValue) Int() int {
	return defaultIntVal
}

func (b uintValue) Int8() int8 {
	return defaultIntVal
}

func (b uintValue) Int16() int16 {
	return defaultIntVal
}

func (b uintValue) Int32() int32 {
	return defaultIntVal
}

func (b uintValue) Int64() int64 {
	return defaultIntVal
}

func (b uintValue) UInt() uint {
	return b.value
}

func (b uintValue) UInt8() uint8 {
	return defaultIntVal
}

func (b uintValue) UInt16() uint16 {
	return defaultIntVal
}

func (b uintValue) UInt32() uint32 {
	return defaultIntVal
}

func (b uintValue) UInt64() uint64 {
	return defaultIntVal
}

func (b uintValue) Byte() byte {
	return defaultIntVal
}

func (b uintValue) Rune() rune {
	return defaultIntVal
}

func (b uintValue) Float32() float32 {
	return defaultFloatVal
}

func (b uintValue) Float64() float64 {
	return defaultFloatVal
}

func (b uintValue) Complex64() complex64 {
	return defaultComplex
}

func (b uintValue) Complex128() complex128 {
	return defaultComplex
}

func (b uintValue) Array() []Value {
	return []Value{}
}

func (b uintValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b uintValue) Strut() Fieldable {
	return defaultFieldable{}
}

func UInt(value uint) Value {
	return uintValue{value}
}
