package main

import tl "github.com/JoelOtter/termloop"
import src "github.com/dev/termFarm/src"

func main() {
  game := tl.NewGame()
  s := src.NewScreen()
  s.LoadMainMenu()
  game.Screen().SetLevel(s.Level)
  game.Start()
}