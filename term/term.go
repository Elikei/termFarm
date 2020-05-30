package term

import tl "github.com/JoelOtter/termloop"
import entity "github.com/dev/termFarm/entity"

func Mainterm() {
	g := tl.NewGame()
	s := entity.NewScreen()
	s.LoadMainMenu()
    g.Screen().SetLevel(s.Level)
	g.Start()

}
