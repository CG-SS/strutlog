package strutlog

import (
	"bytes"
	"fmt"
	"strconv"
)

// marshalJsonValue marshals the given Value into JSON bytes.
// We create a bytes.Buffer and append the Value to it accordingly.
// This function is recursive, since when we unmarshal an Array type, we call marshalJsonValue on every element.
// If the given Value is a Float32Type or Float64Type, floatPrecision and floatFormat will be used to convert the given
// float value into a JSON string.
func marshalJsonValue(value Value, floatPrecision int, floatFormat FloatFormat) []byte {
	if value == nil {
		return nil
	}

	var buffer bytes.Buffer

	switch value.Type() {
	case BoolType:
		buffer.WriteString(strconv.FormatBool(value.Bool()))
	case StringType:
		buffer.WriteByte('"')
		buffer.WriteString(value.String())
		buffer.WriteByte('"')
	case IntType:
		buffer.WriteString(strconv.FormatInt(int64(value.Int()), 10))
	case Int8Type:
		buffer.WriteString(strconv.FormatInt(int64(value.Int8()), 10))
	case Int16Type:
		buffer.WriteString(strconv.FormatInt(int64(value.Int16()), 10))
	case Int32Type:
		buffer.WriteString(strconv.FormatInt(int64(value.Int32()), 10))
	case Int64Type:
		buffer.WriteString(strconv.FormatInt(value.Int64(), 10))
	case UIntType:
		buffer.WriteString(strconv.FormatUint(uint64(value.UInt()), 10))
	case UInt8Type:
		buffer.WriteString(strconv.FormatUint(uint64(value.UInt8()), 10))
	case UInt16Type:
		buffer.WriteString(strconv.FormatUint(uint64(value.UInt16()), 10))
	case UInt32Type:
		buffer.WriteString(strconv.FormatUint(uint64(value.UInt32()), 10))
	case UInt64Type:
		buffer.WriteString(strconv.FormatUint(value.UInt64(), 10))
	case ByteType:
		buffer.WriteString(strconv.FormatUint(uint64(value.Byte()), 10))
	case RuneType:
		buffer.WriteString(strconv.FormatUint(uint64(value.Rune()), 10))
	case Float32Type:
		buffer.WriteString(strconv.FormatFloat(float64(value.Float32()), byte(floatFormat), floatPrecision, 32))
	case Float64Type:
		buffer.WriteString(strconv.FormatFloat(value.Float64(), byte(floatFormat), floatPrecision, 64))
	case Complex64Type:
		buffer.WriteByte('"')
		buffer.WriteString(strconv.FormatComplex(complex128(value.Complex64()), byte(floatFormat), floatPrecision, 64))
		buffer.WriteByte('"')
	case Complex128Type:
		buffer.WriteByte('"')
		buffer.WriteString(strconv.FormatComplex(value.Complex128(), byte(floatFormat), floatPrecision, 128))
		buffer.WriteByte('"')
	case ArrayType:
		buffer.WriteString("[")

		valAr := value.Array()
		fieldsLen := len(valAr)

		for i, v := range valAr {
			buffer.Write(marshalJsonValue(v, floatPrecision, floatFormat))

			if i+1 != fieldsLen {
				buffer.WriteString(",")
			}
		}
		buffer.WriteString("]")
	case MapType:
		buffer.WriteByte('{')

		valuesMap := value.Map()
		valuesMapLen := len(valuesMap)

		i := 0
		for k, v := range valuesMap {
			field := Field{
				Key:   k.String(),
				Value: v,
			}

			buffer.Write(marshalJsonField(field, floatPrecision, floatFormat))

			if i+1 != valuesMapLen {
				buffer.WriteString(",")
			}
			i++
		}

		buffer.WriteByte('}')
	case StrutType:
		buffer.WriteByte('{')

		strutFields := value.Strut().Fields()
		strutFieldsLen := len(strutFields)

		for i, field := range strutFields {
			buffer.Write(marshalJsonField(field, floatPrecision, floatFormat))

			if i+1 != strutFieldsLen {
				buffer.WriteString(",")
			}
		}

		buffer.WriteByte('}')
	}

	return buffer.Bytes()
}

// marshalJsonField marshals the given Field into JSON bytes.
func marshalJsonField(field Field, floatPrecision int, floatFormat FloatFormat) []byte {
	if field.Value == nil || len(field.Key) == 0 {
		return []byte{}
	}

	var buffer bytes.Buffer

	key := field.Key
	buffer.WriteString(fmt.Sprintf("\"%s\":", key))

	value := field.Value
	buffer.Write(marshalJsonValue(value, floatPrecision, floatFormat))

	return buffer.Bytes()
}

// JsonMarshaller is a MessageMarshaller that will build a function in which will marshall a Message into JSON bytes.
func JsonMarshaller(floatPrecision int, floatFormat FloatFormat) MessageMarshaller {
	return func(msg Message) []byte {
		if msg == nil {
			return nil
		}

		var buffer bytes.Buffer
		buffer.WriteString("{")

		fields := msg.Fields()
		fieldsLen := len(fields)
		for i, field := range fields {
			fieldBytes := marshalJsonField(field, floatPrecision, floatFormat)
			n, _ := buffer.Write(fieldBytes)

			if i+1 != fieldsLen && n > 0 {
				buffer.WriteString(",")
			}
		}

		buffer.WriteString("}\n")

		return buffer.Bytes()
	}
}
