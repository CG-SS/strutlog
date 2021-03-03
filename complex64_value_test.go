package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplex64WrapsComplex64Value(t *testing.T) {
	complexVal := complex64(10)
	value := Complex64(complexVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Complex64(), complexVal)
}
