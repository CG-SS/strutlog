package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt16WrapsInt16Value(t *testing.T) {
	intVal := int16(44)
	value := Int16(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Int16(), intVal)
}
