package audio

import (
	"bytes"
	"os"
"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"

)

func Player(p *Playlist, otoCtx *oto.Context, readyChan chan struct {} ) oto.Player {
  fileBytes, err := os.ReadFile("./audio/songs/" + p.CurrentSong())
  if err != nil {
    panic("reading my-file.mp3 failed: " + err.Error())
  }
  
  fileBytesReader := bytes.NewReader(fileBytes)
  decodedMp3, err := mp3.NewDecoder(fileBytesReader)
  if err != nil {
    panic("mp3.NewDecoder failed: " + err.Error())
  }

  <-readyChan
    
  var player *oto.Player
  player = otoCtx.NewPlayer(decodedMp3)
  player.Play()

  go func ()  {
    time_now := time.Now()
    for {
      for player.IsPlaying() {
        time.Sleep(time.Millisecond)
        if int(time.Since(time_now).Round(time.Second)) == int(songDurationMp3(decodedMp3)) * 1000000000 {
          player.Close() 
          keyboard.SimulateKeyPress(keys.Tab)
        }
      }
      if p.Pause == false {
        player.Close()
      }
    }
  }()
  return *player 
  
}
func songDurationMp3(d *mp3.Decoder) int64 {
  samples := d.Length() / 4      // Number of samples.
  audioLength := samples / 44100 // Audio length in seconds.
  return audioLength
}
  
 
  
