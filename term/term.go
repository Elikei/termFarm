package term

import tl "github.com/JoelOtter/termloop"
// import fmt "fmt"

type Pointer struct {
	*tl.Entity
}

func Mainterm() {
	g := tl.NewGame()
	g.Screen().SetFps(60)
	pointer := Pointer{tl.NewEntityFromCanvas(1,4,tl.CanvasFromString(string(" > ")))}

	g.Screen().AddEntity(&pointer)
	g.Screen().AddEntity(tl.NewEntityFromCanvas(5, 4, tl.CanvasFromString(string("1.Start"))))
	g.Screen().AddEntity(tl.NewEntityFromCanvas(5, 5, tl.CanvasFromString(string("2.Load"))))
	g.Screen().AddEntity(tl.NewEntityFromCanvas(5, 6, tl.CanvasFromString(string("3.Quit"))))
	g.Start()
}

func (p *Pointer) Tick(ev tl.Event) {
	_, y:= p.Position()
	if ev.Type == tl.EventKey {
		switch ev.Key {
		case tl.KeyArrowUp:
			if y > 4 {
				p.SetPosition(1, y-1)
			}
		case tl.KeyArrowDown:
			if y < 6 {
				p.SetPosition(1, y+1)
			}
		}
	}
}
