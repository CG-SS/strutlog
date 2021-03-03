package strutlog

type int16Value struct {
	value int16
}

func (b int16Value) Type() PrimitiveType {
	return Int16Type
}

func (b int16Value) Bool() bool {
	return defaultBoolVal
}

func (b int16Value) String() string {
	return defaultStringVal
}

func (b int16Value) Int() int {
	return defaultIntVal
}

func (b int16Value) Int8() int8 {
	return defaultIntVal
}

func (b int16Value) Int16() int16 {
	return b.value
}

func (b int16Value) Int32() int32 {
	return defaultIntVal
}

func (b int16Value) Int64() int64 {
	return defaultIntVal
}

func (b int16Value) UInt() uint {
	return defaultIntVal
}

func (b int16Value) UInt8() uint8 {
	return defaultIntVal
}

func (b int16Value) UInt16() uint16 {
	return defaultIntVal
}

func (b int16Value) UInt32() uint32 {
	return defaultIntVal
}

func (b int16Value) UInt64() uint64 {
	return defaultIntVal
}

func (b int16Value) Byte() byte {
	return defaultIntVal
}

func (b int16Value) Rune() rune {
	return defaultIntVal
}

func (b int16Value) Float32() float32 {
	return defaultFloatVal
}

func (b int16Value) Float64() float64 {
	return defaultFloatVal
}

func (b int16Value) Complex64() complex64 {
	return defaultComplex
}

func (b int16Value) Complex128() complex128 {
	return defaultComplex
}

func (b int16Value) Array() []Value {
	return []Value{}
}

func (b int16Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b int16Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Int16(value int16) Value {
	return int16Value{value}
}
