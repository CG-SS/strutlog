package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolValueWrapsBool(t *testing.T) {
	wrapperValue := Bool(false)
	assert.NotNil(t, wrapperValue)
	assert.Equal(t, false, wrapperValue.Bool())
}
