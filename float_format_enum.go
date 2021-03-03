package strutlog

// FloatFormat is one of:
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise),
// 'G' ('E' for large exponents, 'f' otherwise),
// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
type FloatFormat byte

const (
	// BinaryExponentFloatFormat (-ddddp±ddd, a binary exponent).
	BinaryExponentFloatFormat FloatFormat = 'b'
	// DecimalExponentFloatFormat (-d.dddde±dd, a decimal exponent).
	DecimalExponentFloatFormat FloatFormat = 'e'
	// ExtendedDecimalExponentFloatFormat (-d.ddddE±dd, a decimal exponent).
	ExtendedDecimalExponentFloatFormat FloatFormat = 'E'
	// NoExponentFloatFormat (-ddd.dddd, no exponent).
	NoExponentFloatFormat FloatFormat = 'f'
	// LargeExponentsFloatFormat ('e' for large exponents, 'f' otherwise).
	LargeExponentsFloatFormat FloatFormat = 'g'
	// ExtendedLargeExponentsFloatFormat ('E' for large exponents, 'f' otherwise).
	ExtendedLargeExponentsFloatFormat FloatFormat = 'G'
	// HexFractionBinaryExponentFloatFormat (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent).
	HexFractionBinaryExponentFloatFormat FloatFormat = 'x'
	// ExtendedHexFractionBinaryExponentFloatFormat (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	ExtendedHexFractionBinaryExponentFloatFormat FloatFormat = 'X'
)
