package strutlog

type float64Value struct {
	value float64
}

func (b float64Value) Type() PrimitiveType {
	return Float64Type
}

func (b float64Value) Bool() bool {
	return defaultBoolVal
}

func (b float64Value) String() string {
	return defaultStringVal
}

func (b float64Value) Int() int {
	return defaultIntVal
}

func (b float64Value) Int8() int8 {
	return defaultIntVal
}

func (b float64Value) Int16() int16 {
	return defaultIntVal
}

func (b float64Value) Int32() int32 {
	return defaultIntVal
}

func (b float64Value) Int64() int64 {
	return defaultIntVal
}

func (b float64Value) UInt() uint {
	return defaultIntVal
}

func (b float64Value) UInt8() uint8 {
	return defaultIntVal
}

func (b float64Value) UInt16() uint16 {
	return defaultIntVal
}

func (b float64Value) UInt32() uint32 {
	return defaultIntVal
}

func (b float64Value) UInt64() uint64 {
	return defaultIntVal
}

func (b float64Value) Byte() byte {
	return defaultIntVal
}

func (b float64Value) Rune() rune {
	return defaultIntVal
}

func (b float64Value) Float32() float32 {
	return defaultFloatVal
}

func (b float64Value) Float64() float64 {
	return b.value
}

func (b float64Value) Complex64() complex64 {
	return defaultComplex
}

func (b float64Value) Complex128() complex128 {
	return defaultComplex
}

func (b float64Value) Array() []Value {
	return []Value{}
}

func (b float64Value) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b float64Value) Strut() Fieldable {
	return defaultFieldable{}
}

func Float64(value float64) Value {
	return float64Value{value}
}
