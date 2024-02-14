package pm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPm(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		cmd := New(true, true, true, true, 8, "*")
		out, err := cmd.Gen()
		assert.NoError(t, err)
		fmt.Println(out)
	})
	t.Run("t2_min12", func(t *testing.T) {
		cmd := NewDefaultMin12()
		out, err := cmd.Gen()
		assert.NoError(t, err)
		fmt.Println(out)
	})
}
