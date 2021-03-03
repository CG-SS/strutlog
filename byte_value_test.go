package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteValueWrapsByte(t *testing.T) {
	byteVal := byte(1)
	value := Byte(byteVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Byte(), byteVal)
}
