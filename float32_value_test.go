package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat32WrapsFloat32Value(t *testing.T) {
	floatVal := float32(10.5)
	value := Float32(floatVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Float32(), floatVal)
}
