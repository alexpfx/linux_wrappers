package rofi

import (
	"fmt"
	"github.com/alexpfx/linux_wrappers/linux"
	"github.com/bitfield/script"
	"strings"
)

const (
	rofiCmd             = "rofi"
	rofiPrompt          = "-p"
	rofiAutoSelect      = "-auto-select"
	rofiThemeStr        = "-theme-str"
	rofiFormat          = "-format"
	rofiMode            = "-show"
	rofiDmenu           = "-dmenu"
	rofiErrMsg          = "-e"
	rofiSep             = "-sep"
	rofiSelect          = "-select"
	rofiCaseInsensitive = "-i"
	rofiPangoMarkup     = "-markup-rows"
	rofiMatching        = "-matching"
)

var kbMatrix = [4][11]rune{
	{
		'1', '2', '3', '4', '5', ' ', '6', '7', '8', '9', '0',
	},
	{
		'q', 'w', 'e', 'r', 't', ' ', 'y', 'u', 'I', 'o', 'p',
	},
	{
		'a', 's', 'd', 'f', 'g', ' ', 'h', 'j', 'k', 'l', 'ç',
	},
	{
		'z', 'x', 'c', 'v', 'b', ' ', 'n', 'm', ',', '.', ';',
	},
}

type rofiMenu struct {
	args      string
	actionMap map[rune]linux.KeyAction
}

func NewDMenu(prompt string) linux.Rofi {
	b := builder{
		dMenu:  true,
		prompt: prompt,
	}
	args := b.buildArgs()
	return rofiMenu{
		args: args,
	}
}

func NewMessageMenu(errMsg string) linux.Rofi {
	b := builder{
		errMessage: errMsg,
	}
	return rofiMenu{
		args: b.buildArgs(),
	}

}

func NewKeyboardMenu(actionMap map[rune]linux.KeyAction) linux.RofiKeyboard {
	b := builder{
		autoSelect:      true,
		dMenu:           true,
		caseInsensitive: true,
		sep:             "|",
		format:          "s",
		pangoMarkup:     true,
		matching:        "prefix",
		themeStr: `
element {
    padding: 2px ;
    cursor:  pointer;
    spacing: 5px ;
    border:  1;
	children: [ element-text ];

}
listview {
    padding:      2px 2px 2px ;
    scrollbar:    false;
    border-color: var(separatorcolor);
    spacing:      5px ;
    fixed-height: 0;
    border:       1px dash 1px 1px ;
    columns: 11 ;
    fixed-columns: true ;
    lines:  4;
    flow: horizontal;
}
element-text {
    background-color: transparent;
    cursor:           inherit;
    highlight:        inherit;
    text-color:       inherit;
    vertical-align: 0.5;
    horizontal-align: 0.1;
}
`,
	}

	return rofiMenu{
		args:      b.buildArgs(),
		actionMap: actionMap,
	}
}

func (r rofiMenu) Show() (string, error) {
	var p *script.Pipe
	cmdStr := fmt.Sprintf("%s %s", rofiCmd, r.args)

	menuStr := ""

	for i := 0; i < len(kbMatrix); i++ {
		for j := 0; j < len(kbMatrix[i]); j++ {
			key := kbMatrix[i][j]
			if act, ok := r.actionMap[key]; ok {
				menuStr += fmt.Sprintf("<b>%s</b>: <small><i>%s</i></small>\000display\x1f%s|", strings.ToUpper(string(key)), act.Label, string(key))
			} else if key != ' ' {
				menuStr += fmt.Sprintf("<b>%s</b>:   |", strings.ToUpper(string(key)))
			} else {
				menuStr += " |"
			}
		}
	}

	p = script.Echo(menuStr).Exec(cmdStr)
	return p.String()

}

func (r rofiMenu) Run(input string) (string, error) {
	var p *script.Pipe
	cmdStr := fmt.Sprintf("%s %s", rofiCmd, r.args)
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
	//'s' selected string
	//'i' index (0 - (N-1))
	//'d' index (1 - N)
	//'q' quote string
	//'p' Selected string stripped from Pango markup (Needs to be a valid string)
	//'f' filter string (user action)
	//'F' quoted filter string (user action)

	format          string
	mode            string
	dMenu           bool
	errMessage      string
	sep             string
	rselect         string
	caseInsensitive bool
	pangoMarkup     bool
	matching        string
	//normal: match the int string
	//regex: match a regex input
	//glob: match a glob pattern
	//fuzzy: do a fuzzy match
	//prefix: match prefix

}

func (r builder) buildArgs() string {
	argSlice := make([]string, 0)

	argSlice = linux.AppendIf(argSlice, rofiMode, r.mode)
	argSlice = linux.AppendIf(argSlice, rofiDmenu, r.dMenu)

	argSlice = linux.AppendIf(argSlice, rofiPrompt, r.prompt)
	argSlice = linux.AppendIf(argSlice, rofiAutoSelect, r.autoSelect)

	argSlice = linux.AppendIf(argSlice, rofiThemeStr, r.themeStr)
	argSlice = linux.AppendIf(argSlice, rofiFormat, r.format)
	argSlice = linux.AppendIf(argSlice, rofiErrMsg, r.errMessage)
	argSlice = linux.AppendIf(argSlice, rofiSep, r.sep)
	argSlice = linux.AppendIf(argSlice, rofiCaseInsensitive, r.caseInsensitive)
	argSlice = linux.AppendIf(argSlice, rofiPangoMarkup, r.pangoMarkup)
	argSlice = linux.AppendIf(argSlice, rofiSelect, r.rselect)
	argSlice = linux.AppendIf(argSlice, rofiMatching, r.matching)
	return strings.Join(argSlice, " ")

}
