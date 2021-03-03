package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64WrapsInt64Value(t *testing.T) {
	intVal := int64(44)
	value := Int64(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Int64(), intVal)
}
