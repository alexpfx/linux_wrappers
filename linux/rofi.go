package linux

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"strings"
)

func NewDMenu(prompt string) Rofi {
	b := builder{
		dMenu:  true,
		prompt: prompt,
	}
	args := b.buildArgs()
	return rofiMenu{
		args: args,
	}
}

func NewMessageMenu(errMsg string) Rofi {
	b := builder{
		errMessage: errMsg,
	}
	return rofiMenu{
		args: b.buildArgs(),
	}

}

func (r rofiMenu) Run(input string) (string, error) {
	var p *script.Pipe
	cmdStr := fmt.Sprintf("%s %s", rofiCmd, strings.Join(r.args, " "))
	log.Printf("cmd: %s", cmdStr)
	if input == "" {
		p = script.Exec(cmdStr)
		return p.String()
	}
	p = script.Echo(input).Exec(cmdStr)
	return p.String()
}

type builder struct {
	prompt     string
	autoSelect bool
	themeStr   string
	//	//'s' selected string
	//	//'i' index (0 - (N-1))
	//	//'d' index (1 - N)
	//	//'q' quote string
	//	//'p' Selected string stripped from Pango markup (Needs to be a valid string)
	//	//'f' filter string (user action)
	//	//'F' quoted filter string (user action)
	format     string
	mode       string
	dMenu      bool
	errMessage string
}

func (r builder) buildArgs() []string {
	argSlice := make([]string, 0)

	argSlice = appendIf(argSlice, rofiMode, r.mode)
	argSlice = appendIf(argSlice, rofiDmenu, r.dMenu)

	argSlice = appendIf(argSlice, rofiPrompt, r.prompt)
	argSlice = appendIf(argSlice, rofiAutoSelect, r.autoSelect)

	argSlice = appendIf(argSlice, rofiThemeStr, r.themeStr)
	argSlice = appendIf(argSlice, rofiFormat, r.format)
	argSlice = appendIf(argSlice, rofiErrMsg, r.errMessage)

	return argSlice

}
