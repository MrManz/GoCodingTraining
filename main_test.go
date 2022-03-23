package GoCodingTraining

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUnique(t *testing.T) {
	assert.Equal(t, true, hasUniqueCharacters("abcd"))
	assert.Equal(t, true, hasUniqueCharacters("a"))
	assert.Equal(t, false, hasUniqueCharacters("abca"))
	assert.Equal(t, false, hasUniqueCharacters("aabb"))
	assert.Equal(t, false, hasUniqueCharacters("cc"))
}
