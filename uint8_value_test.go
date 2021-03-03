package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUInt8WrapsUInt8Value(t *testing.T) {
	intVal := uint8(12)
	value := UInt8(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.UInt8(), intVal)
}
