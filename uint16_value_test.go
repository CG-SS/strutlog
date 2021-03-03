package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUInt16WrapsUInt16Value(t *testing.T) {
	intVal := uint16(16)
	value := UInt16(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.UInt16(), intVal)
}
