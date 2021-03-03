package strutlog

type intValue struct {
	value int
}

func (b intValue) Type() PrimitiveType {
	return IntType
}

func (b intValue) Bool() bool {
	return defaultBoolVal
}

func (b intValue) String() string {
	return defaultStringVal
}

func (b intValue) Int() int {
	return b.value
}

func (b intValue) Int8() int8 {
	return defaultIntVal
}

func (b intValue) Int16() int16 {
	return defaultIntVal
}

func (b intValue) Int32() int32 {
	return defaultIntVal
}

func (b intValue) Int64() int64 {
	return defaultIntVal
}

func (b intValue) UInt() uint {
	return defaultIntVal
}

func (b intValue) UInt8() uint8 {
	return defaultIntVal
}

func (b intValue) UInt16() uint16 {
	return defaultIntVal
}

func (b intValue) UInt32() uint32 {
	return defaultIntVal
}

func (b intValue) UInt64() uint64 {
	return defaultIntVal
}

func (b intValue) Byte() byte {
	return defaultIntVal
}

func (b intValue) Rune() rune {
	return defaultIntVal
}

func (b intValue) Float32() float32 {
	return defaultFloatVal
}

func (b intValue) Float64() float64 {
	return defaultFloatVal
}

func (b intValue) Complex64() complex64 {
	return defaultComplex
}

func (b intValue) Complex128() complex128 {
	return defaultComplex
}

func (b intValue) Array() []Value {
	return []Value{}
}

func (b intValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b intValue) Strut() Fieldable {
	return defaultFieldable{}
}

func Int(value int) Value {
	return intValue{value}
}
