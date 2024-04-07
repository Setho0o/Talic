package main

import (
	"github.com/Setho0o/Talic/audio"
)

func main() {
  go audio.Player("./audio/songs/Doomsday.flac")
  Tui()
  Keys() 
}
func Esc() {
  Options()
}
