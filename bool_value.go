package strutlog

type boolValue struct {
	value bool
}

func (b boolValue) Type() PrimitiveType {
	return BoolType
}

func (b boolValue) Bool() bool {
	return b.value
}

func (b boolValue) String() string {
	return defaultStringVal
}

func (b boolValue) Int() int {
	return defaultIntVal
}

func (b boolValue) Int8() int8 {
	return defaultIntVal
}

func (b boolValue) Int16() int16 {
	return defaultIntVal
}

func (b boolValue) Int32() int32 {
	return defaultIntVal
}

func (b boolValue) Int64() int64 {
	return defaultIntVal
}

func (b boolValue) UInt() uint {
	return defaultIntVal
}

func (b boolValue) UInt8() uint8 {
	return defaultIntVal
}

func (b boolValue) UInt16() uint16 {
	return defaultIntVal
}

func (b boolValue) UInt32() uint32 {
	return defaultIntVal
}

func (b boolValue) UInt64() uint64 {
	return defaultIntVal
}

func (b boolValue) Byte() byte {
	return defaultIntVal
}

func (b boolValue) Rune() rune {
	return defaultIntVal
}

func (b boolValue) Float32() float32 {
	return defaultFloatVal
}

func (b boolValue) Float64() float64 {
	return defaultFloatVal
}

func (b boolValue) Complex64() complex64 {
	return defaultComplex
}

func (b boolValue) Complex128() complex128 {
	return defaultComplex
}

func (b boolValue) Array() []Value {
	return []Value{}
}

func (b boolValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b boolValue) Strut() Fieldable {
	return defaultFieldable{}
}

func Bool(value bool) Value {
	return boolValue{value}
}
