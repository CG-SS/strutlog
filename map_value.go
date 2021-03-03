package strutlog

type mapValue struct {
	value map[Value]Value
}

func (b mapValue) Type() PrimitiveType {
	return MapType
}

func (b mapValue) Bool() bool {
	return defaultBoolVal
}

func (b mapValue) String() string {
	return defaultStringVal
}

func (b mapValue) Int() int {
	return defaultIntVal
}

func (b mapValue) Int8() int8 {
	return defaultIntVal
}

func (b mapValue) Int16() int16 {
	return defaultIntVal
}

func (b mapValue) Int32() int32 {
	return defaultIntVal
}

func (b mapValue) Int64() int64 {
	return defaultIntVal
}

func (b mapValue) UInt() uint {
	return defaultIntVal
}

func (b mapValue) UInt8() uint8 {
	return defaultIntVal
}

func (b mapValue) UInt16() uint16 {
	return defaultIntVal
}

func (b mapValue) UInt32() uint32 {
	return defaultIntVal
}

func (b mapValue) UInt64() uint64 {
	return defaultIntVal
}

func (b mapValue) Byte() byte {
	return defaultIntVal
}

func (b mapValue) Rune() rune {
	return defaultIntVal
}

func (b mapValue) Float32() float32 {
	return defaultFloatVal
}

func (b mapValue) Float64() float64 {
	return defaultFloatVal
}

func (b mapValue) Complex64() complex64 {
	return defaultComplex
}

func (b mapValue) Complex128() complex128 {
	return defaultComplex
}

func (b mapValue) Array() []Value {
	return []Value{}
}

func (b mapValue) Map() map[Value]Value {
	return b.value
}

func (b mapValue) Strut() Fieldable {
	return defaultFieldable{}
}

func Map(value map[Value]Value) Value {
	return mapValue{value}
}
