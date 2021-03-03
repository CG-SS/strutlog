package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrutWrapsStrutValue(t *testing.T) {
	strutVal := defaultFieldable{}
	value := Strut(strutVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Strut(), strutVal)
}
