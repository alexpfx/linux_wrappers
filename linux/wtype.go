package linux

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"strings"
)

const wtypeCmd = "wtype"
const wtypePressMod = "-M"
const wtypeReleaseMod = "-m"
const wtypeKey = "-k"
const wtypeDelayBetweenKeyStrokes = "-d"
const wtypeDelayBeforeKeyStrokes = "-s"
const wtypePressKey = "-P"
const wtypeReleaseKey = "-p"

type wtype struct {
	args string
}

func (w wtype) Run(text string) (string, error) {
	cmdStr := fmt.Sprintf(`%s %s '%s'`, wtypeCmd, w.args, text)
	log.Printf("cmd: %s", cmdStr)
	p := script.Exec(cmdStr)
	return p.String()
}

func NewWType(builder WTypeBuilder) WType {
	args := builder.buildArgs()
	return wtype{
		args: args,
	}
}

func Run() {
}

type WTypeBuilder struct {
	PressModifier          string
	ReleaseModifier        string
	PressKey               string
	ReleaseKey             string
	Type                   string
	DelayBetweenKeyStrokes string
	DelayBeforeKeyStrokes  string
}

func (r WTypeBuilder) buildArgs() string {
	argSlice := make([]string, 0)

	argSlice = AppendIf(argSlice, wtypePressMod, r.PressModifier)
	argSlice = AppendIf(argSlice, wtypeReleaseMod, r.ReleaseModifier)
	argSlice = AppendIf(argSlice, wtypePressKey, r.PressKey)
	argSlice = AppendIf(argSlice, wtypeReleaseKey, r.ReleaseKey)
	argSlice = AppendIf(argSlice, wtypeKey, r.Type)

	argSlice = AppendIf(argSlice, wtypeDelayBetweenKeyStrokes, r.DelayBetweenKeyStrokes)
	argSlice = AppendIf(argSlice, wtypeDelayBeforeKeyStrokes, r.DelayBeforeKeyStrokes)

	return strings.Join(argSlice, " ")
}

//
//Modifier can be one of "shift", "capslock", "ctrl", "logo", "win", "alt", "altgr". Beware that the modifiers get released automatically once the program terminates.
//
//Named keys are resolved by libxkbcommon, valid identifiers include "Left" and "Home".
//
//-M MOD
//Press modifier MOD.
//-m MOD
//Release modifier MOD.
//-P KEY
//Press key KEY.
//-p KEY
//Release key KEY.
//-k KEY
//Type (press and release) key KEY.
//-d TIME
//Sleep for TIME milliseconds between keystrokes when typing texts. Can be used multiple times, default 0.
//-s TIME
//Sleep for TIME milliseconds before interpreting the following options. This can be used to perform more complicated modifier sequences.
//-
//Read text to type from stdin. This option can appear only once.
