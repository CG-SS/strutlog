package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayValueWrapsArray(t *testing.T) {
	intVal1 := intValue{1}
	intVal2 := intValue{2}

	values := []Value{
		intVal1,
		intVal2,
	}

	wrappedArray := Array(values)
	assert.NotNil(t, wrappedArray)

	array := wrappedArray.Array()
	assert.NotNil(t, array)
	assert.Len(t, array, 2)
	assert.Equal(t, array[0], intVal1)
	assert.Equal(t, array[1], intVal2)
}
