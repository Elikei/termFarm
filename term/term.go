package term

import tl "github.com/JoelOtter/termloop"

type Pointer struct {
	*tl.Entity
}

func Mainterm() {
	g := tl.NewGame()
	g.Screen().SetFps(60)
	// pointer := Pointer{tl.NewEntity(1,1,1,1)}
	// pointer.SetCell(0, 0, &tl.Cell({Fg: tl.ColorRed, Ch: '>'}))
	pointer := Pointer{tl.NewEntityFromCanvas(1,1,tl.CanvasFromString(string(" > ")))}
	e1 := tl.NewEntityFromCanvas(5, 1, tl.CanvasFromString(string("1.Start")))
	e2 := tl.NewEntityFromCanvas(5, 2, tl.CanvasFromString(string("2.Option")))
	e3 := tl.NewEntityFromCanvas(5, 3, tl.CanvasFromString(string("3.Quit")))
	g.Screen().AddEntity(&pointer)
	g.Screen().AddEntity(e1)
	g.Screen().AddEntity(e2)
	g.Screen().AddEntity(e3)
	g.Start()
}

func (p *Pointer) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey {
		switch ev.Key {
		case tl.KeyArrowUp:
			p.SetPosition(1,3)
		case tl.KeyArrowDown:
			p.SetPosition(1,2)
		}
	}
}