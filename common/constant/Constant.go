package constant

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

var (
	ScreenWidth  = 1280
	ScreenHeight = 720
	GameMode     = ModeTitle
)
