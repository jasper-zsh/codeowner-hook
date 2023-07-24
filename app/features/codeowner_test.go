package features

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchPattern(t *testing.T) {
	file := "app/foo/bar/baz.go"
	trueRules := []string{
		"app/foo/* @owner1 @owner2",
		"app/*/bar @owner1 @owner2",
		"app/foo/bar/baz.go @owner1 @owner2",
		"app/foo @owner1 @owner2",
	}
	falseRules := []string{
		"app/baz @owner1 @owner2",
		"app/*/baz/* @owner1 @owner2",
		"app/**/bar.go @owner1 @owner2",
		"app/foo/bar/bar.go @owner1 @owner2",
	}
	for _, r := range trueRules {
		rule := NewCodeOwnerRuleFromLine(r)
		assert.True(t, rule.Match(file), r)
	}
	for _, r := range falseRules {
		rule := NewCodeOwnerRuleFromLine(r)
		assert.False(t, rule.Match(file), r)
	}
}
