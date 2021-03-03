package strutlog

type int8Value struct {
	value int8
}

func (b int8Value) Type() PrimitiveType {
	return Int8Type
}

func (b int8Value) Bool() bool {
	return defaultBoolVal
}

func (b int8Value) String() string {
	return defaultStringVal
}

func (b int8Value) Int() int {
	return defaultIntVal
}

func (b int8Value) Int8() int8 {
	return b.value
}

func (b int8Value) Int16() int16 {
	return defaultIntVal
}

func (b int8Value) Int32() int32 {
	return defaultIntVal
}

func (b int8Value) Int64() int64 {
	return defaultIntVal
}

func (b int8Value) UInt() uint {
	return defaultIntVal
}

func (b int8Value) UInt8() uint8 {
	return defaultIntVal
}

func (b int8Value) UInt16() uint16 {
	return defaultIntVal
}

func (b int8Value) UInt32() uint32 {
	return defaultIntVal
}

func (b int8Value) UInt64() uint64 {
	return defaultIntVal
}

func (b int8Value) Byte() byte {
	return defaultIntVal
}

func (b int8Value) Rune() rune {
	return defaultIntVal
}

func (b int8Value) Float32() float32 {
	return defaultFloatVal
}

func (b int8Value) Float64() float64 {
	return defaultFloatVal
}

func (b int8Value) Complex64() complex64 {
	return defaultComplex
}

func (b int8Value) Complex128() complex128 {
	return defaultComplex
}

func (b int8Value) Array() []Value {
	return []Value{}
}

func (b int8Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b int8Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Int8(value int8) Value {
	return int8Value{value}
}
