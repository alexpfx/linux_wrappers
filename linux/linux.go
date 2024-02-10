package linux

import (
	"fmt"
	"log"
	"reflect"
)

const (
	rofiCmd        = "rofi"
	rofiPrompt     = "-p"
	rofiAutoSelect = "-auto-select"
	rofiThemeStr   = "-theme-str"
	rofiFormat     = "-format"
	rofiMode       = "-show"
	rofiDmenu      = "-dmenu"
	rofiErrMsg     = "-e"
)

type Rofi interface {
	Run(input string) (string, error)
}
type rofiMenu struct {
	args []string
}

type WType interface {
	Run(text string) (string, error)
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
