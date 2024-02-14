package rofi

type KeyAction struct {
	Label  string
	Action func() string
}

type RofiKeyboard interface {
	DMenu() (string, error)
}

type Rofi interface {
	ShowDMenu(input string) (string, error)
}

type Response struct {
	Output string
	Err    error
}
