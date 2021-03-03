package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntWrapsIntValue(t *testing.T) {
	intVal := 44
	value := Int(intVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Int(), intVal)
}
