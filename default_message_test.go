package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultMessageAppendsBool(t *testing.T) {
	key := "test"
	message := defaultMessage{}
	message2 := message.Bool(key, true)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, BoolType, field.Value.Type())
	assert.Equal(t, true, field.Value.Bool())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsString(t *testing.T) {
	key := "test"
	value := "test val"
	message := defaultMessage{}
	message2 := message.String(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, StringType, field.Value.Type())
	assert.Equal(t, value, field.Value.String())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsInt(t *testing.T) {
	key := "test"
	value := 1
	message := defaultMessage{}
	message2 := message.Int(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, IntType, field.Value.Type())
	assert.Equal(t, value, field.Value.Int())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsInt8(t *testing.T) {
	key := "test"
	value := int8(1)
	message := defaultMessage{}
	message2 := message.Int8(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Int8Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Int8())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsInt16(t *testing.T) {
	key := "test"
	value := int16(1)
	message := defaultMessage{}
	message2 := message.Int16(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Int16Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Int16())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsInt32(t *testing.T) {
	key := "test"
	value := int32(12)
	message := defaultMessage{}
	message2 := message.Int32(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Int32Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Int32())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsInt64(t *testing.T) {
	key := "test"
	value := int64(12)
	message := defaultMessage{}
	message2 := message.Int64(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Int64Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Int64())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsUInt(t *testing.T) {
	key := "test"
	value := uint(1)
	message := defaultMessage{}
	message2 := message.UInt(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, UIntType, field.Value.Type())
	assert.Equal(t, value, field.Value.UInt())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsUInt8(t *testing.T) {
	key := "test"
	value := uint8(1)
	message := defaultMessage{}
	message2 := message.UInt8(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, UInt8Type, field.Value.Type())
	assert.Equal(t, value, field.Value.UInt8())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsUInt16(t *testing.T) {
	key := "test"
	value := uint16(1)
	message := defaultMessage{}
	message2 := message.UInt16(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, UInt16Type, field.Value.Type())
	assert.Equal(t, value, field.Value.UInt16())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsUInt32(t *testing.T) {
	key := "test"
	value := uint32(12)
	message := defaultMessage{}
	message2 := message.UInt32(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, UInt32Type, field.Value.Type())
	assert.Equal(t, value, field.Value.UInt32())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsUInt64(t *testing.T) {
	key := "test"
	value := uint64(12)
	message := defaultMessage{}
	message2 := message.UInt64(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, UInt64Type, field.Value.Type())
	assert.Equal(t, value, field.Value.UInt64())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsByte(t *testing.T) {
	key := "test"
	value := byte(5)
	message := defaultMessage{}
	message2 := message.Byte(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, ByteType, field.Value.Type())
	assert.Equal(t, value, field.Value.Byte())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsRune(t *testing.T) {
	key := "test"
	value := rune(5)
	message := defaultMessage{}
	message2 := message.Rune(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, RuneType, field.Value.Type())
	assert.Equal(t, value, field.Value.Rune())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsFloat32(t *testing.T) {
	key := "test"
	value := float32(5.6)
	message := defaultMessage{}
	message2 := message.Float32(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Float32Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Float32())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsFloat64(t *testing.T) {
	key := "test"
	value := 5.6
	message := defaultMessage{}
	message2 := message.Float64(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Float64Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Float64())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsComplex64(t *testing.T) {
	key := "test"
	value := complex64(88)
	message := defaultMessage{}
	message2 := message.Complex64(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Complex64Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Complex64())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsComplex128(t *testing.T) {
	key := "test"
	value := complex128(88)
	message := defaultMessage{}
	message2 := message.Complex128(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, Complex128Type, field.Value.Type())
	assert.Equal(t, value, field.Value.Complex128())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsArray(t *testing.T) {
	key := "array test"
	value := []Value{
		intValue{12},
		boolValue{true},
		float64Value{7.8},
	}

	message := defaultMessage{}
	message2 := message.Array(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, ArrayType, field.Value.Type())
	assert.Equal(t, value, field.Value.Array())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsMap(t *testing.T) {
	key := "map test"
	value := map[Value]Value{
		intValue{1}: boolValue{true},
	}

	message := defaultMessage{}
	message2 := message.Map(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, MapType, field.Value.Type())
	assert.Equal(t, value, field.Value.Map())
	assert.Equal(t, key, field.Key)
}

func TestDefaultMessageAppendsStrut(t *testing.T) {
	key := "strut test"
	value := defaultFieldable{}

	message := defaultMessage{}
	message2 := message.Strut(key, value)

	// We make sure that the original message wasn't modified.
	assert.Len(t, message.Fields(), 0)

	assert.Len(t, message2.Fields(), 1)
	field := message2.Fields()[0]
	assert.Equal(t, StrutType, field.Value.Type())
	assert.Equal(t, value, field.Value.Strut())
	assert.Equal(t, key, field.Key)
}
