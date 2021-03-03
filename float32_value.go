package strutlog

type float32Value struct {
	value float32
}

func (b float32Value) Type() PrimitiveType {
	return Float32Type
}

func (b float32Value) Bool() bool {
	return defaultBoolVal
}

func (b float32Value) String() string {
	return defaultStringVal
}

func (b float32Value) Int() int {
	return defaultIntVal
}

func (b float32Value) Int8() int8 {
	return defaultIntVal
}

func (b float32Value) Int16() int16 {
	return defaultIntVal
}

func (b float32Value) Int32() int32 {
	return defaultIntVal
}

func (b float32Value) Int64() int64 {
	return defaultIntVal
}

func (b float32Value) UInt() uint {
	return defaultIntVal
}

func (b float32Value) UInt8() uint8 {
	return defaultIntVal
}

func (b float32Value) UInt16() uint16 {
	return defaultIntVal
}

func (b float32Value) UInt32() uint32 {
	return defaultIntVal
}

func (b float32Value) UInt64() uint64 {
	return defaultIntVal
}

func (b float32Value) Byte() byte {
	return defaultIntVal
}

func (b float32Value) Rune() rune {
	return defaultIntVal
}

func (b float32Value) Float32() float32 {
	return b.value
}

func (b float32Value) Float64() float64 {
	return defaultFloatVal
}

func (b float32Value) Complex64() complex64 {
	return defaultComplex
}

func (b float32Value) Complex128() complex128 {
	return defaultComplex
}

func (b float32Value) Array() []Value {
	return []Value{}
}

func (b float32Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b float32Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Float32(value float32) Value {
	return float32Value{value}
}
