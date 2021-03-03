package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRuneWrapsRuneValue(t *testing.T) {
	runeVal := rune(6)
	value := Rune(runeVal)
	assert.NotNil(t, value)
	assert.Equal(t, value.Rune(), runeVal)
}
