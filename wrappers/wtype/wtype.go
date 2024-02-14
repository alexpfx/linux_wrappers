package wtype

import (
	"fmt"
	"github.com/alexpfx/linux_wrappers/wrappers"
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

type WType interface {
	ShowDMenu(text string) (string, error)
}

type wtype struct {
	args string
}

func (w wtype) ShowDMenu(text string) (string, error) {
	cmdStr := fmt.Sprintf(`%s %s '%s'`, wtypeCmd, w.args, text)
	log.Printf("cmd: %s", cmdStr)
	p := script.Exec(cmdStr)
	return p.String()
}

func New(builder Builder) WType {
	args := builder.buildArgs()
	return wtype{
		args: args,
	}
}

type Builder struct {
	PressModifier          string
	ReleaseModifier        string
	PressKey               string
	ReleaseKey             string
	Type                   string
	DelayBetweenKeyStrokes string
	DelayBeforeKeyStrokes  string
}

func (r Builder) buildArgs() string {
	argSlice := make([]string, 0)

	argSlice = wrappers.AppendIf(argSlice, wtypePressMod, r.PressModifier)
	argSlice = wrappers.AppendIf(argSlice, wtypeReleaseMod, r.ReleaseModifier)
	argSlice = wrappers.AppendIf(argSlice, wtypePressKey, r.PressKey)
	argSlice = wrappers.AppendIf(argSlice, wtypeReleaseKey, r.ReleaseKey)
	argSlice = wrappers.AppendIf(argSlice, wtypeKey, r.Type)

	argSlice = wrappers.AppendIf(argSlice, wtypeDelayBetweenKeyStrokes, r.DelayBetweenKeyStrokes)
	argSlice = wrappers.AppendIf(argSlice, wtypeDelayBeforeKeyStrokes, r.DelayBeforeKeyStrokes)

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
