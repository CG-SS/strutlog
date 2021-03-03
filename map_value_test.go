package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapWrapsMapValue(t *testing.T) {
	var mapVal map[Value]Value
	value := Map(mapVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Map(), mapVal)
}
