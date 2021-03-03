package strutlog

type runeValue struct {
	value rune
}

func (b runeValue) Type() PrimitiveType {
	return RuneType
}

func (b runeValue) Bool() bool {
	return defaultBoolVal
}

func (b runeValue) String() string {
	return defaultStringVal
}

func (b runeValue) Int() int {
	return defaultIntVal
}

func (b runeValue) Int8() int8 {
	return defaultIntVal
}

func (b runeValue) Int16() int16 {
	return defaultIntVal
}

func (b runeValue) Int32() int32 {
	return defaultIntVal
}

func (b runeValue) Int64() int64 {
	return defaultIntVal
}

func (b runeValue) UInt() uint {
	return defaultIntVal
}

func (b runeValue) UInt8() uint8 {
	return defaultIntVal
}

func (b runeValue) UInt16() uint16 {
	return defaultIntVal
}

func (b runeValue) UInt32() uint32 {
	return defaultIntVal
}

func (b runeValue) UInt64() uint64 {
	return defaultIntVal
}

func (b runeValue) Byte() byte {
	return defaultIntVal
}

func (b runeValue) Rune() rune {
	return b.value
}

func (b runeValue) Float32() float32 {
	return defaultFloatVal
}

func (b runeValue) Float64() float64 {
	return defaultFloatVal
}

func (b runeValue) Complex64() complex64 {
	return defaultComplex
}

func (b runeValue) Complex128() complex128 {
	return defaultComplex
}

func (b runeValue) Array() []Value {
	return []Value{}
}

func (b runeValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b runeValue) Strut() Fieldable {
	return defaultFieldable{}
}

func Rune(value rune) Value {
	return runeValue{value}
}
