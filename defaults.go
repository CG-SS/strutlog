package strutlog

const (
	// defaultBoolVal is the default value used for Bool when the PrimitiveType is not BoolType.
	defaultBoolVal = false
	// defaultStringVal is the default string value used for String when the PrimitiveType is not StringType.
	defaultStringVal = ""
	// defaultStringVal is the default string value used for all the Int types when their corresponding PrimitiveType
	// are not correct.
	defaultIntVal = 0
	// defaultFloatVal is the default float value used for all the Float types when their corresponding PrimitiveType
	// are not correct.
	defaultFloatVal = 0.0
	// defaultComplex is the default complex value used for all the Complex types when their corresponding PrimitiveType
	// are not correct.
	defaultComplex = complex(0, 0)
)

// defaultFieldable is a struct that implements the Fieldable interface and is the default value used for Strut when the
// PrimitiveType is not StrutType.
type defaultFieldable struct {
}

// Fields will return an empty Field array.
func (d defaultFieldable) Fields() []Field {
	return []Field{}
}
