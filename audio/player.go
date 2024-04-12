package audio

import (
	"bytes"
	"fmt"
	"os"
	"time"

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
    
    player := otoCtx.NewPlayer(decodedMp3)
    player.Play()
    go func ()  {
      for player.IsPlaying() {
        time.Sleep(time.Millisecond)
      }
      Next(p, *player)
    }()
    return *player 
  
}
func Close(player oto.Player) {
  player.Close()
}
func Pause(player oto.Player) {
  player.Pause()
}
func Next(p *Playlist, player oto.Player) {
  p.Song++
  p.OpenFalse()
  fmt.Println(p.Open)
  Close(player)
}
  

  
  
