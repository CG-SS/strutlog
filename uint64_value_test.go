package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUInt64WrapsUInt64Value(t *testing.T) {
	intVal := uint64(16)
	value := UInt64(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.UInt64(), intVal)
}
