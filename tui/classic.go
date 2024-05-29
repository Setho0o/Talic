package tui

import (
  "os"
  "fmt"

	gt "github.com/buger/goterm"
	"github.com/muesli/termenv"

  "github.com/Setho0o/Talic/utils"
)
func Classic() {
  output := termenv.NewOutput(os.Stdout) 
  output.HideCursor()
  output.AltScreen()
  boxLeft()
  boxUp()
  boxDown()
}

func boxLeft() {
  box := gt.NewBox(50|gt.PCT , gt.Height(), 0)
  box.Border = "━ ┃ ┏ ┓ ┗ ┛ "

	songs, _ := utils.GetSongs()
  for _, e := range songs {
    fmt.Fprint(box, e, "\n") 
  }

  gt.Print(box.String())
  gt.Flush()
}

func boxUp() { 
  box := gt.NewBox(50|gt.PCT , 52|gt.PCT , 0)
  box.Border = "━ ┃ ┏ ┓ ┗ ┛ "

  gt.Print(gt.MoveTo(box.String(), 52|gt.PCT, 5|gt.PCT))
  gt.Flush()

}

func boxDown() { 
  box := gt.NewBox(50|gt.PCT , 51|gt.PCT , 0)
  box.Border = "━ ┃ ┏ ┓ ┗ ┛ "
	
  gt.Print(gt.MoveTo(box.String(), 52|gt.PCT, 55|gt.PCT)) 
  gt.Flush()

}

