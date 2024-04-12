package main

import (
	"os"
	"os/exec"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"github.com/Setho0o/Talic/audio"
	"github.com/ebitengine/oto/v3"
)
func Keys(p *audio.Playlist, player oto.Player) {
  keyboard.Listen(func(key keys.Key) (stop bool, err error) {
    if !p.Open { //Determins if the playlist has been closed ex - changing a song 
      return true, nil
    }
    if key.Code == keys.RuneKey {
      switch key.String() {
      case "q" :
        cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
        cmd.Run()
        return true, nil 
      }
    } else {
        switch key.Code {
        case keys.CtrlC, keys.Escape:
          Esc()
        case keys.Tab:
         // Kind of confusing the way I designed it but when a new song needs to be played we just end this process and go on to the next cycle. The song is updated in the player file 
          audio.Next(p, player)  
          return true, nil 
          
      }
    }
    return false, nil 
  })
}
