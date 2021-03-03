package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64WrapsFloat64Value(t *testing.T) {
	floatVal := 10.5
	value := Float64(floatVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Float64(), floatVal)
}
