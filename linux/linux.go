package linux

const (
	rofiCmd        = "rofi"
	rofiPrompt     = "-p"
	rofiAutoSelect = "-auto-select"
	rofiThemeStr   = "-theme-str"
	rofiFormat     = "-format"
	rofiMode       = "-show"
	rofiDmenu      = "-dmenu"
)

type Rofi interface {
	Run(input string) (string, error)
}
type rofiMenu struct {
	args []string
}
