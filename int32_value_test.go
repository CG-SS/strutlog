package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt32WrapsInt32Value(t *testing.T) {
	intVal := int32(44)
	value := Int32(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Int32(), intVal)
}
