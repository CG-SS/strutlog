package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt8WrapsInt8Value(t *testing.T) {
	intVal := int8(44)
	value := Int8(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Int8(), intVal)
}
