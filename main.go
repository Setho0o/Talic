package main

import (
	"github.com/ebitengine/oto/v3"

	"github.com/Setho0o/Talic/audio"
	"github.com/Setho0o/Talic/keys"
	"github.com/Setho0o/Talic/tui"
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

  playlist := audio.Playlist{}
  for {
   player := audio.Player(&playlist, otoCtx, readyChan)
    tui.Minimal(playlist)
    end := keys.Keys(&playlist, &player)
    if end == true {
      break
    }
  }  

}
