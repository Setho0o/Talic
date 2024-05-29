package keys

import (
	"os"
	"os/exec"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"github.com/Setho0o/Talic/audio"
	"github.com/ebitengine/oto/v3"
)
func Keys(p *audio.Playlist, player *oto.Player) bool {
  end := false
  keyboard.Listen(func(key keys.Key) (stop bool, err error) {
    if key.Code == keys.RuneKey {
      switch key.String() {
      case "q" :
        cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
        cmd.Run()
        end = true
        return true, nil
      case "a" : 
        player.Close() 
        time.Sleep(time.Millisecond * 100)
        p.PreviousPlaylist()
        p.Song = 0 
        return true, nil
      case "d" : 
        player.Close() 
        time.Sleep(time.Millisecond * 100)
        p.NextPlaylist()
        p.Song = 0 
        return true, nil

      }
    } else {
      switch key.Code {

      case keys.CtrlC, keys.Escape:

      case keys.Tab, keys.Right:   
        player.Close() 
        time.Sleep(time.Millisecond * 100)
        p.NextSong()
        return true, nil

      case keys.Space: 
        if p.Pause == false {
          player.Pause()
          p.Pause = true 
        } else {
          player.Play()
          p.Pause = false 
        }
      case keys.Up:
        player.SetVolume(player.Volume() + 0.10)
        if player.Volume() > 1 {
          player.SetVolume(1)
        }
      case keys.Down:
        player.SetVolume(player.Volume() - 0.10)
        if player.Volume() < 0 {
          player.SetVolume(0)
        }
      case keys.Left:
        player.Close() 
        time.Sleep(time.Millisecond * 100)
        p.PreviousSong()
        return true, nil
      }
    }
    return false, nil 
  })
  return end
}
