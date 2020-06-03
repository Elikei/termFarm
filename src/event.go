package src

import os "os"

func NewProfileEvent() {
  NewProfile()
}

func LoadProfileEvent() {
  LoadProfile()
}

func ExitEvent() {
  os.Exit(1)
}
