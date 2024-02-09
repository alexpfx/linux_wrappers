package linux

import (
	"fmt"

	"github.com/bitfield/script"
)

func (w WType) Type(text string, delayMs int) (string, error) {
	return script.Exec(fmt.Sprintf("wtype -d %d '%s'", delayMs, text)).String()
}

type WType struct {
}
