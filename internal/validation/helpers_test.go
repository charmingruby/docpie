package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	emptyValue := ""
	assert.True(t, IsEmpty(emptyValue))

	notEmptyValue := "not empty"
	assert.False(t, IsEmpty(notEmptyValue))
}

func TestIsLower(t *testing.T) {
	word := "four"

	equalResult := IsLower(word, len(word))
	assert.False(t, equalResult)

	lowerResult := IsLower(word, len(word)+1)
	assert.True(t, lowerResult)

	greaterResult := IsLower(word, len(word)-1)
	assert.False(t, greaterResult)
}

func TestIsLowerOrEqual(t *testing.T) {
	word := "four"

	equalResult := IsLowerOrEqual(word, len(word))
	assert.True(t, equalResult)

	lowerResult := IsLowerOrEqual(word, len(word)+1)
	assert.True(t, lowerResult)

	greaterResult := IsLowerOrEqual(word, len(word)-1)
	assert.False(t, greaterResult)
}

func TestIsGreater(t *testing.T) {
	word := "four"

	equalResult := IsGreater(word, len(word))
	assert.False(t, equalResult)

	lowerResult := IsGreater(word, len(word)+1)
	assert.False(t, lowerResult)

	greaterResult := IsGreater(word, len(word)-1)
	assert.True(t, greaterResult)
}

func TestIsGreaterOrEqual(t *testing.T) {
	word := "four"

	equalResult := IsGreaterOrEqual(word, len(word))
	assert.True(t, equalResult)

	lowerResult := IsGreaterOrEqual(word, len(word)+1)
	assert.False(t, lowerResult)

	greaterResult := IsGreaterOrEqual(word, len(word)-1)
	assert.True(t, greaterResult)
}

func TestIsEmail(t *testing.T) {
	validEmail := "john@example.com"
	assert.True(t, IsEmail(validEmail))

	invalidEmail := "bad example"
	assert.False(t, IsEmail(invalidEmail))

}
