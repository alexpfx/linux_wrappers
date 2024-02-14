package wtype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWType(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		w := New(Builder{
			PressModifier:         "",
			ReleaseModifier:       "",
			PressKey:              "",
			ReleaseKey:            "",
			Type:                  "",
			DelayBeforeKeyStrokes: "1100",
		})

		r, err := w.ShowDMenu("texto$")
		assert.NoError(t, err, r)

	})

}
