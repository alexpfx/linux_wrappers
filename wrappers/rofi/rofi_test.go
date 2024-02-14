package rofi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRofi(t *testing.T) {
	type args struct {
		prompt string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "t1", args: struct{ prompt string }{prompt: "Mensagem de prompt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rofi := New(tt.args.prompt)
			_, err := rofi.ShowDMenu("teste")
			assert.NoError(t, err, "Nao esperava erro")
		})
	}
}

func TestNewMessageMenu(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		rofi := NewMessage("mensagem de erro")
		_, err := rofi.ShowDMenu("")
		assert.NoError(t, err, "Nao esperava erro")
	})
}

func TestNewKeyboardMenu(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		rofi := NewKeyboardMenu(map[rune]KeyAction{
			'a': {
				Label: "Hora",
				Action: func() string {
					return "hora"
				},
			},
			'm': {
				Label: "Massa",
				Action: func() string {
					return "massa"
				},
			},
		})
		out, err := rofi.Show()
		fmt.Println(out)
		assert.NoError(t, err, "Nao esperava erro")
	})
}
