package linux

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"reflect"
	"strings"
)

func NewDMenu(prompt string) Rofi {
	b := builder{
		dMenu:  true,
		format: "f",
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

func appendIf(res []string, argName string, pValue interface{}) []string {
	if pValue == nil {
		return res
	}
	vs := reflect.ValueOf(pValue)

	switch vs.Kind() {
	case reflect.String:
		if vs.String() != "" {
			res = appendArgName(res, argName)
			res = append(res, fmt.Sprintf("'%s'", vs.String()))
			return res
		}

	case reflect.Bool:
		if vs.Bool() {
			res = appendArgName(res, argName)
			return res
		}
	case reflect.Slice:
		for i := 0; i < vs.Len(); i++ {
			vIndex := vs.Index(i)

			if i == 0 {
				res = appendIf(res, argName, vIndex.Interface())
			} else {
				res = appendIf(res, "", vIndex.Interface())
			}
		}
	case reflect.Int:
		log.Fatal(fmt.Errorf("cannot be int, use string instead"))
	}

	return res
}

func appendArgName(slice []string, argName string) []string {
	if argName == "" {
		return slice
	}
	slice = append(slice, argName)
	return slice
}
