package strutlog

type strutValue struct {
	value Fieldable
}

func (b strutValue) Type() PrimitiveType {
	return StrutType
}

func (b strutValue) Bool() bool {
	return defaultBoolVal
}

func (b strutValue) String() string {
	return defaultStringVal
}

func (b strutValue) Int() int {
	return defaultIntVal
}

func (b strutValue) Int8() int8 {
	return defaultIntVal
}

func (b strutValue) Int16() int16 {
	return defaultIntVal
}

func (b strutValue) Int32() int32 {
	return defaultIntVal
}

func (b strutValue) Int64() int64 {
	return defaultIntVal
}

func (b strutValue) UInt() uint {
	return defaultIntVal
}

func (b strutValue) UInt8() uint8 {
	return defaultIntVal
}

func (b strutValue) UInt16() uint16 {
	return defaultIntVal
}

func (b strutValue) UInt32() uint32 {
	return defaultIntVal
}

func (b strutValue) UInt64() uint64 {
	return defaultIntVal
}

func (b strutValue) Byte() byte {
	return defaultIntVal
}

func (b strutValue) Rune() rune {
	return defaultIntVal
}

func (b strutValue) Float32() float32 {
	return defaultFloatVal
}

func (b strutValue) Float64() float64 {
	return defaultFloatVal
}

func (b strutValue) Complex64() complex64 {
	return defaultComplex
}

func (b strutValue) Complex128() complex128 {
	return defaultComplex
}

func (b strutValue) Array() []Value {
	return []Value{}
}

func (b strutValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b strutValue) Strut() Fieldable {
	return b.value
}

func Strut(value Fieldable) Value {
	return strutValue{value}
}
