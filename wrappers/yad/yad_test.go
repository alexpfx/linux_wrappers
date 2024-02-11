package yad

import (
	"testing"
)

func TestNewButtonBar(t *testing.T) {
	t.Run("t0", func(t *testing.T) {
		bb := NewButtonBar(nil)
		bb.Show()
	})
}
