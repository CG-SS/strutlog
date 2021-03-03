package strutlog

type PrimitiveType string

const (
	BoolType       PrimitiveType = "bool"
	StringType     PrimitiveType = "string"
	IntType        PrimitiveType = "int"
	Int8Type       PrimitiveType = "int8"
	Int16Type      PrimitiveType = "int16"
	Int32Type      PrimitiveType = "int32"
	Int64Type      PrimitiveType = "int64"
	UIntType       PrimitiveType = "uint"
	UInt8Type      PrimitiveType = "uint8"
	UInt16Type     PrimitiveType = "uint16"
	UInt32Type     PrimitiveType = "uint32"
	UInt64Type     PrimitiveType = "uint64"
	ByteType       PrimitiveType = "byte"
	RuneType       PrimitiveType = "rune"
	Float32Type    PrimitiveType = "float32"
	Float64Type    PrimitiveType = "float64"
	Complex64Type  PrimitiveType = "complex64"
	Complex128Type PrimitiveType = "complex128"
	ArrayType      PrimitiveType = "array"
	MapType        PrimitiveType = "map"
	StrutType      PrimitiveType = "strut"
)
