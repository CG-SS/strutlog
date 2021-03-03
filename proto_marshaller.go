package strutlog

import (
	"google.golang.org/protobuf/proto"
	"strconv"
	"strutlog/generated"
)

// protoMarshalValue marshals a given Value to a protobuf generated.Value type.
func protoMarshalValue(value Value) *generated.Value {
	switch value.Type() {
	case BoolType:
		return &generated.Value{
			Type:  generated.PrimitiveType_BoolType,
			Value: &generated.Value_BoolValue{BoolValue: value.Bool()},
		}
	case StringType:
		return &generated.Value{
			Type:  generated.PrimitiveType_StringType,
			Value: &generated.Value_StringValue{StringValue: value.String()},
		}
	case IntType:
		return &generated.Value{
			Type:  generated.PrimitiveType_IntType,
			Value: &generated.Value_Int32Value{Int32Value: int32(value.Int())},
		}
	case Int8Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_Int8Type,
			Value: &generated.Value_Int32Value{Int32Value: int32(value.Int8())},
		}
	case Int16Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_Int16Type,
			Value: &generated.Value_Int32Value{Int32Value: int32(value.Int16())},
		}
	case Int32Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_Int32Type,
			Value: &generated.Value_Int32Value{Int32Value: value.Int32()},
		}
	case Int64Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_Int64Type,
			Value: &generated.Value_Int64Value{Int64Value: value.Int64()},
		}
	case UIntType:
		return &generated.Value{
			Type:  generated.PrimitiveType_UIntType,
			Value: &generated.Value_Uint32Value{Uint32Value: uint32(value.UInt())},
		}
	case UInt8Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_UInt8Type,
			Value: &generated.Value_Uint32Value{Uint32Value: uint32(value.UInt8())},
		}
	case UInt16Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_UInt16Type,
			Value: &generated.Value_Uint32Value{Uint32Value: uint32(value.UInt16())},
		}
	case UInt32Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_UInt32Type,
			Value: &generated.Value_Uint32Value{Uint32Value: value.UInt32()},
		}
	case UInt64Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_UInt64Type,
			Value: &generated.Value_Int64Value{Int64Value: value.Int64()},
		}
	case ByteType:
		return &generated.Value{
			Type:  generated.PrimitiveType_ByteType,
			Value: &generated.Value_Uint32Value{Uint32Value: uint32(value.Byte())},
		}
	case RuneType:
		return &generated.Value{
			Type:  generated.PrimitiveType_RuneType,
			Value: &generated.Value_Int32Value{Int32Value: value.Rune()},
		}
	case Float32Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_Float32Type,
			Value: &generated.Value_Float32Value{Float32Value: value.Float32()},
		}
	case Float64Type:
		return &generated.Value{
			Type:  generated.PrimitiveType_Float64Type,
			Value: &generated.Value_Float64Value{Float64Value: value.Float64()},
		}
	case Complex64Type:
		complex64Fmt := strconv.FormatComplex(complex128(value.Complex64()), 'f', -1, 64)

		return &generated.Value{
			Type:  generated.PrimitiveType_Complex64Type,
			Value: &generated.Value_StringValue{StringValue: complex64Fmt},
		}
	case Complex128Type:
		complex128Fmt := strconv.FormatComplex(value.Complex128(), 'f', -1, 128)

		return &generated.Value{
			Type:  generated.PrimitiveType_Complex128Type,
			Value: &generated.Value_StringValue{StringValue: complex128Fmt},
		}
	case ArrayType:
		values := value.Array()
		generatedValues := make([]*generated.Value, 0, len(values))

		for _, v := range values {
			generatedValue := protoMarshalValue(v)
			generatedValues = append(generatedValues, generatedValue)
		}

		return &generated.Value{
			Type:       generated.PrimitiveType_ArrayType,
			ArrayValue: generatedValues,
		}
	case MapType:
		valuesMap := value.Map()
		generatedMap := make(map[string]*generated.Value, len(valuesMap))

		for k, v := range valuesMap {
			generatedValue := protoMarshalValue(v)

			generatedMap[k.String()] = generatedValue
		}

		return &generated.Value{
			Type:     generated.PrimitiveType_MapType,
			MapValue: generatedMap,
		}
	case StrutType:
		strutFields := value.Strut().Fields()
		generatedMap := make(map[string]*generated.Value, len(strutFields))

		for _, f := range strutFields {
			generatedValue := protoMarshalValue(f.Value)

			generatedMap[f.Key] = generatedValue
		}

		return &generated.Value{
			Type:     generated.PrimitiveType_StrutType,
			MapValue: generatedMap,
		}
	}

	return nil
}

// protoMarshalField marshals a given Field to a protobuf generated.Field type.
func protoMarshalField(field Field) *generated.Field {
	generatedValue := protoMarshalValue(field.Value)

	return &generated.Field{
		Key:   field.Key,
		Value: generatedValue,
	}
}

// ProtoMarshaller is a MessageMarshaller that marshals a given Message to protobuf.
func ProtoMarshaller() MessageMarshaller {
	return func(message Message) []byte {
		if message == nil {
			return []byte{}
		}

		fields := message.Fields()
		generatedFields := make([]*generated.Field, 0, len(fields))

		for _, f := range fields {
			generatedField := protoMarshalField(f)
			generatedFields = append(generatedFields, generatedField)
		}

		generatedMessage := &generated.Message{Fields: generatedFields}
		marshal, err := proto.Marshal(generatedMessage)
		if err != nil {
			return []byte{}
		}

		return marshal
	}
}
