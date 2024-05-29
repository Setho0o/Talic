package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/Setho0o/Talic/audio"
	"github.com/Setho0o/Talic/utils"
	color "github.com/TwiN/go-color"
	gt "github.com/buger/goterm"
	"github.com/muesli/termenv"
)
func Minimal(playlist audio.Playlist)  {
  utils.Clear()
  
  box := gt.NewBox(100|gt.PCT , 100|gt.PCT , 0)
  box.Border = "     "
  box.PaddingX = 2
  
  p := strings.Trim(playlist.CurrentPlaylist(),"audio/music/")
  s := playlist.GetSongs()
  fmt.Fprint(box, color.Colorize(color.Yellow, color.InBold(p)), "\n\n")
  for i, e := range s {
    sym := "◇"
    if i == playlist.Song { 
      sym = "◆"
    } 
    e = color.Colorize(color.Green, e)
    fmt.Fprint(box, sym+" "+e,"\n", gt.Color(strings.Repeat("─",len(e) + 1), gt.BLUE),"\n")
  }

  output := termenv.NewOutput(os.Stdout)
  output.HideCursor()
  gt.Print(box.String())
  gt.Flush()
}
