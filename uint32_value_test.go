package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUInt32WrapsUInt32Value(t *testing.T) {
	intVal := uint32(16)
	value := UInt32(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.UInt32(), intVal)
}
