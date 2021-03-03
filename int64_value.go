package strutlog

type int64Value struct {
	value int64
}

func (b int64Value) Type() PrimitiveType {
	return Int64Type
}

func (b int64Value) Bool() bool {
	return defaultBoolVal
}

func (b int64Value) String() string {
	return defaultStringVal
}

func (b int64Value) Int() int {
	return defaultIntVal
}

func (b int64Value) Int8() int8 {
	return defaultIntVal
}

func (b int64Value) Int16() int16 {
	return defaultIntVal
}

func (b int64Value) Int32() int32 {
	return defaultIntVal
}

func (b int64Value) Int64() int64 {
	return b.value
}

func (b int64Value) UInt() uint {
	return defaultIntVal
}

func (b int64Value) UInt8() uint8 {
	return defaultIntVal
}

func (b int64Value) UInt16() uint16 {
	return defaultIntVal
}

func (b int64Value) UInt32() uint32 {
	return defaultIntVal
}

func (b int64Value) UInt64() uint64 {
	return defaultIntVal
}

func (b int64Value) Byte() byte {
	return defaultIntVal
}

func (b int64Value) Rune() rune {
	return defaultIntVal
}

func (b int64Value) Float32() float32 {
	return defaultFloatVal
}

func (b int64Value) Float64() float64 {
	return defaultFloatVal
}

func (b int64Value) Complex64() complex64 {
	return defaultComplex
}

func (b int64Value) Complex128() complex128 {
	return defaultComplex
}

func (b int64Value) Array() []Value {
	return []Value{}
}

func (b int64Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b int64Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Int64(value int64) Value {
	return int64Value{value}
}
