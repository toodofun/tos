package application

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

type Theme string

type App struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Page       string `json:"page"`
	Width      uint   `json:"width"`
	Height     uint   `json:"height"`
	X          uint   `json:"x"`
	Y          uint   `json:"y"`
	Theme      Theme  `json:"theme"`
	Background string `json:"background"`
	Singleton  bool   `json:"singleton"`
}
