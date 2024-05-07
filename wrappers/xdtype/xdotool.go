package xdtype

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"strings"
)

// Constantes para os comandos do xdotool
const (
	xdotoolCmd             = "xdotool"
	xdotoolKey             = "key"
	xdotoolType            = "type"
	xdotoolKeyDown         = "keydown"
	xdotoolKeyUp           = "keyup"
	xdotoolDelay           = "--delay"
)

// XDoTool interface para interagir com o xdotool
type XDoTool interface {
	Type(text string) (string, error)
	KeyPress(key string) (string, error)
	KeyRelease(key string) (string, error)
}

// xdotool struct que implementa a interface XDoTool
type xdotool struct {
	args string
}

// Type digita um texto usando o comando 'type' do xdotool
func (x xdotool) Type(text string) (string, error) {
	cmdStr := fmt.Sprintf("%s %s %s '%s'", xdotoolCmd, xdotoolType, x.args, text)
	log.Printf("Type command: %s", cmdStr)
	p := script.Exec(cmdStr)
	return p.String()
}

// KeyPress pressiona uma tecla usando o comando 'keydown' do xdotool
func (x xdotool) KeyPress(key string) (string, error) {
	cmdStr := fmt.Sprintf("%s %s %s", xdotoolCmd, xdotoolKeyDown, key)
	log.Printf("KeyPress command: %s", cmdStr)
	p := script.Exec(cmdStr)
	return p.String()
}

// KeyRelease solta uma tecla usando o comando 'keyup' do xdotool
func (x xdotool) KeyRelease(key string) (string, error) {
	cmdStr := fmt.Sprintf("%s %s %s", xdotoolCmd, xdotoolKeyUp, key)
	log.Printf("KeyRelease command: %s", cmdStr)
	p := script.Exec(cmdStr)
	return p.String()
}

// New cria uma nova instância do xdotool com argumentos personalizados
func New(builder Builder) XDoTool {
	args := builder.buildArgs()
	return xdotool{
		args: args,
	}
}

// Builder para construir argumentos personalizados para o xdotool
type Builder struct {
	Delay string
}

// buildArgs constrói a string de argumentos para ser usada com o xdotool
func (b Builder) buildArgs() string {
	argSlice := make([]string, 0)

	if b.Delay != "" {
		argSlice = append(argSlice, fmt.Sprintf("%s %s", xdotoolDelay, b.Delay))
	}

	return strings.Join(argSlice, " ")
}
