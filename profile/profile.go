package profile

import os "os"
import fmt "fmt"
import random "math/rand"
// import term "github.com/dev/termFarm/term"

func NewProfile() {
  profileName := profileNameGen()
  fmt.Println("\n\n Welcome, the owner of " + profileName)
  file, err := os.Create(profileName)
  defer file.Close()
  if err != nil {
  	fmt.Println("Create profile failed: [%s]", err)
  	os.Exit(1)
  } else {
  	_, err := file.Write([]byte("SAVE"))
  	if err != nil {
  		fmt.Println("Something goes wrong: [%s]", err)
  		os.Exit(1)
  	}
  }
  term.Mainterm()
}

func LoadProfile() {

}

func profileNameGen() string {
  firstName,lastName := [2]string{"Blue", "Sweet"}, [2]string{"Home", "Land"}
  return firstName[random.Intn(len(firstName))] + lastName[random.Intn(len(lastName))]
}