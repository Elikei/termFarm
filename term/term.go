package term

import tl "github.com/JoelOtter/termloop"

func Mainterm() {
	g := tl.NewGame()
	s := NewScreen()
	s.LoadMainMenu()
    g.Screen().SetLevel(s.Level)
	g.Start()

}
