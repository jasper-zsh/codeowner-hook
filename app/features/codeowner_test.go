package features

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchPattern(t *testing.T) {
	file := "app/foo/bar/baz.go"
	assert.True(t, matchPatternLevels(file, []string{"app", "foo", "*"}))
	assert.True(t, matchPatternLevels(file, []string{"app", "*", "bar"}))
	assert.True(t, matchPatternLevels(file, []string{"app", "foo", "bar", "baz.go"}))
	assert.True(t, matchPatternLevels(file, []string{"app", "foo"}))
	assert.False(t, matchPatternLevels(file, []string{"app", "baz"}))
	assert.False(t, matchPatternLevels(file, []string{"app", "*", "baz", "*"}))
	assert.False(t, matchPatternLevels(file, []string{"app", "**", "bar.go"}))
	assert.False(t, matchPatternLevels(file, []string{"app", "foo", "bar", "bar.go"}))
}
