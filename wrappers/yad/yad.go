package yad

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"strings"
)

type Yad interface {
	DMenu() (string, error)
}

type yad struct {
	args string
}

func (r yad) DMenu() (string, error) {
	cmdStr := fmt.Sprintf("%s %s", "yad", r.args)
	log.Println(cmdStr)
	pipe := script.Exec(cmdStr)
	out, err := pipe.String()

	return strings.TrimSpace(out), err
}
