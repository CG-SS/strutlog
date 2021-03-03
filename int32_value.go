package strutlog

type int32Value struct {
	value int32
}

func (b int32Value) Type() PrimitiveType {
	return Int32Type
}

func (b int32Value) Bool() bool {
	return defaultBoolVal
}

func (b int32Value) String() string {
	return defaultStringVal
}

func (b int32Value) Int() int {
	return defaultIntVal
}

func (b int32Value) Int8() int8 {
	return defaultIntVal
}

func (b int32Value) Int16() int16 {
	return defaultIntVal
}

func (b int32Value) Int32() int32 {
	return b.value
}

func (b int32Value) Int64() int64 {
	return defaultIntVal
}

func (b int32Value) UInt() uint {
	return defaultIntVal
}

func (b int32Value) UInt8() uint8 {
	return defaultIntVal
}

func (b int32Value) UInt16() uint16 {
	return defaultIntVal
}

func (b int32Value) UInt32() uint32 {
	return defaultIntVal
}

func (b int32Value) UInt64() uint64 {
	return defaultIntVal
}

func (b int32Value) Byte() byte {
	return defaultIntVal
}

func (b int32Value) Rune() rune {
	return defaultIntVal
}

func (b int32Value) Float32() float32 {
	return defaultFloatVal
}

func (b int32Value) Float64() float64 {
	return defaultFloatVal
}

func (b int32Value) Complex64() complex64 {
	return defaultComplex
}

func (b int32Value) Complex128() complex128 {
	return defaultComplex
}

func (b int32Value) Array() []Value {
	return []Value{}
}

func (b int32Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b int32Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Int32(value int32) Value {
	return int32Value{value}
}
