package yad

import (
	"fmt"
	"github.com/alexpfx/linux_wrappers/wrappers"
	"github.com/bitfield/script"
	"log"
	"strings"
)

type Btn struct {
	Label string
	Key   rune
}

var (
	lines = []string{
		"1234567890",
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm<>;",
	}
)

func NewButtonBar(keymap []Btn) wrappers.Yad {
	btns := make([]string, 0)
	for _, keys := range lines {
		for _, key := range keys {
			if key == ' ' {
				continue
			}
			btns = append(btns, fmt.Sprintf(`--field=" _%c:BTN"`, key))
		}
	}
	return yad{args: fmt.Sprintf("--form %s --columns 10 --output-by-row", strings.Join(btns, " "))}
}

type yad struct {
	args string
}

func (r yad) Show() (string, error) {
	cmdStr := fmt.Sprintf("%s %s", "yad", r.args)
	log.Println(cmdStr)
	pipe := script.Exec(cmdStr)
	out, err := pipe.String()

	return strings.TrimSpace(out), err
}
