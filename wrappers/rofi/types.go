package rofi

type KeyAction struct {
	Label  string
	Action func() string
}

type KeyboardLayout interface {
	Show() (string, error)
}

type Rofi interface {
	ShowDMenu(input string) (string, error)
}
