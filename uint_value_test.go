package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUIntWrapsUIntValue(t *testing.T) {
	intVal := uint(16)
	value := UInt(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.UInt(), intVal)
}
