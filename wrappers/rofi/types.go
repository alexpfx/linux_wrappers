package rofi

type KeyAction struct {
	Label  string
	Action func() string
}

type RofiKeyboard interface {
	Show() (string, error)
}

type Rofi interface {
	Run(input string) (string, error)
}

type Response struct {
	Output string
	Err    error
}
