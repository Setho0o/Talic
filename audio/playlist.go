package audio

import (

	"github.com/Setho0o/Talic/utils"
)

type Playlist struct {
  Song int
  Pause bool
}

func (p *Playlist) CurrentSong() string {
  songs := utils.GetSongs()
  return songs[p.Song]
}

func (p *Playlist) NextSong() {
  if p.Song == len(utils.GetSongs()) - 1  { 
    p.Song = 0
  } else {
    p.Song++
  }
}
func (p *Playlist) PreviousSong() {
  if p.Song <= 0  { 
    p.Song = len(utils.GetSongs()) - 1
  } else {
    p.Song--
  }
}



