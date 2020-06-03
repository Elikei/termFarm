package src

import tl "github.com/JoelOtter/termloop"

type Screen struct {
	Level      *tl.BaseLevel
	EntityList []*tl.Entity
	pointer    *Pointer
}

type Pointer struct {
	*tl.Entity
}

func NewScreen() *Screen {
	s := new(Screen)
	s.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorDefault, Fg: tl.ColorDefault})
	s.EntityList = make([]*tl.Entity, 20)
	pointer := Pointer{tl.NewEntityFromCanvas(1, 4, tl.CanvasFromString(string(" > ")))}
	s.Level.AddEntity(&pointer)
	return s
}

func (s *Screen) GetAllEntity() []*tl.Entity {
	return s.EntityList
}

func (s *Screen) ClearEntity() {
	for index, _ := range s.EntityList {
		s.Level.RemoveEntity(s.EntityList[index])
		s.EntityList[index] = nil
	}
}

func (s *Screen) LoadMainMenu() {

	s.addEntity(tl.NewEntityFromCanvas(5, 4, tl.CanvasFromString(string("1.New Game"))))
	s.addEntity(tl.NewEntityFromCanvas(5, 5, tl.CanvasFromString(string("2.Load"))))
	s.addEntity(tl.NewEntityFromCanvas(5, 6, tl.CanvasFromString(string("3.Quit"))))
}

// func (s *Screen) loadFarm() {

// }

func (s *Screen) addEntity(e *tl.Entity) {
	s.Level.AddEntity(e)
	s.EntityList = append(s.EntityList, e)
}

func (p *Pointer) Tick(ev tl.Event) {
	_, y := p.Position()
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
		case tl.KeyEnter:
			mainMenuRoute(y)
		}
	}
}

func mainMenuRoute(position int) {
	switch position {
	case 4:
		NewProfileEvent()
	case 5:
		LoadProfileEvent()
	default:
		ExitEvent()
	}
}
