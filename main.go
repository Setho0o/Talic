package main

import (
	"fmt"

  "github.com/ebitengine/oto/v3"
	"github.com/Setho0o/Talic/audio"
)

func main() {
  op := &oto.NewContextOptions{}
  op.SampleRate = 44100
  op.ChannelCount = 2
  op.Format = oto.FormatSignedInt16LE
    
  otoCtx, readyChan, err := oto.NewContext(op)
    
  if err != nil {
    panic("oto.NewContext failed: " + err.Error())
  }

  playlist := audio.Playlist{Song: 0, Open: true}
  for {
    playlist.OpenTrue()
    player := audio.Player(&playlist, otoCtx, readyChan)
    Tui()
    Keys(&playlist, player)
    fmt.Println(playlist.Open, playlist.Song)
    
    if playlist.Open {
      break
    }
  } 
}


func Esc() {
  Options()
}
