package strutlog

type byteValue struct {
	value byte
}

func (b byteValue) Type() PrimitiveType {
	return ByteType
}

func (b byteValue) Bool() bool {
	return defaultBoolVal
}

func (b byteValue) String() string {
	return defaultStringVal
}

func (b byteValue) Int() int {
	return defaultIntVal
}

func (b byteValue) Int8() int8 {
	return defaultIntVal
}

func (b byteValue) Int16() int16 {
	return defaultIntVal
}

func (b byteValue) Int32() int32 {
	return defaultIntVal
}

func (b byteValue) Int64() int64 {
	return defaultIntVal
}

func (b byteValue) UInt() uint {
	return defaultIntVal
}

func (b byteValue) UInt8() uint8 {
	return defaultIntVal
}

func (b byteValue) UInt16() uint16 {
	return defaultIntVal
}

func (b byteValue) UInt32() uint32 {
	return defaultIntVal
}

func (b byteValue) UInt64() uint64 {
	return defaultIntVal
}

func (b byteValue) Byte() byte {
	return b.value
}

func (b byteValue) Rune() rune {
	return defaultIntVal
}

func (b byteValue) Float32() float32 {
	return defaultFloatVal
}

func (b byteValue) Float64() float64 {
	return defaultFloatVal
}

func (b byteValue) Complex64() complex64 {
	return defaultComplex
}

func (b byteValue) Complex128() complex128 {
	return defaultComplex
}

func (b byteValue) Array() []Value {
	return []Value{}
}

func (b byteValue) Map() map[Value]Value {
	return map[Value]Value{}
}

func (b byteValue) Strut() Fieldable {
	return defaultFieldable{}
}

func Byte(value byte) Value {
	return byteValue{value}
}
