package chapter3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	t.Run("has suffix", func(t *testing.T) {
		got := HasSuffix("theapple", "apple")
		assert.True(t, got)
	})
}
