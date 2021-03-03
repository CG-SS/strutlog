package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringWrapsStringValue(t *testing.T) {
	stringVal := "test"
	value := String(stringVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.String(), stringVal)
}
