package event

import profile "github.com/dev/termFarm/profile"
import os "os"
// import term "github.com/dev/termFarm/term"

func NewProfileEvent() {
  profile.NewProfile()
}

func LoadProfileEvent() {
  profile.LoadProfile()
}

func ExitEvent() {
  os.Exit(1)
}

type trans interface {
	NewP()
}