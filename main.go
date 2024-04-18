package main

import (
//	"fmt"

	"github.com/ebitengine/oto/v3"

	"github.com/Setho0o/Talic/audio"
	//"github.com/Setho0o/Talic/utils"
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
    Tui()

    end := Keys(&playlist, &player)
  //  fmt.Println(utils.GetSongs(), "\n", playlist.Song)
    if end == true {
      break
    }
  }    
}
