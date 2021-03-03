package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplex128WrapsComplex128Value(t *testing.T) {
	complexVal := complex128(12)
	value := Complex128(complexVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Complex128(), complexVal)
}
