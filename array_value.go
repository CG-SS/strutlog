package strutlog

type arrayValue struct {
	value []Value
}

func (b arrayValue) Type() PrimitiveType {
	return ArrayType
}

func (b arrayValue) Bool() bool {
	return defaultBoolVal
}

func (b arrayValue) String() string {
	return defaultStringVal
}

func (b arrayValue) Int() int {
	return defaultIntVal
}

func (b arrayValue) Int8() int8 {
	return defaultIntVal
}

func (b arrayValue) Int16() int16 {
	return defaultIntVal
}

func (b arrayValue) Int32() int32 {
	return defaultIntVal
}

func (b arrayValue) Int64() int64 {
	return defaultIntVal
}

func (b arrayValue) UInt() uint {
	return defaultIntVal
}

func (b arrayValue) UInt8() uint8 {
	return defaultIntVal
}

func (b arrayValue) UInt16() uint16 {
	return defaultIntVal
}

func (b arrayValue) UInt32() uint32 {
	return defaultIntVal
}

func (b arrayValue) UInt64() uint64 {
	return defaultIntVal
}

func (b arrayValue) Byte() byte {
	return defaultIntVal
}

func (b arrayValue) Rune() rune {
	return defaultIntVal
}

func (b arrayValue) Float32() float32 {
	return defaultFloatVal
}

func (b arrayValue) Float64() float64 {
	return defaultFloatVal
}

func (b arrayValue) Complex64() complex64 {
	return defaultComplex
}

func (b arrayValue) Complex128() complex128 {
	return defaultComplex
}

func (b arrayValue) Array() []Value {
	return b.value
}

func (b arrayValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b arrayValue) Strut() Fieldable {
	return defaultFieldable{}
}

func Array(value []Value) Value {
	return arrayValue{value}
}
