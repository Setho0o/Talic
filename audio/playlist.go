package audio

import (

	"github.com/Setho0o/Talic/utils"
)

type Playlist struct {
  Song int
  Open bool
}

func (p Playlist) CurrentSong() string {
  songs := utils.GetSongs()
  return songs[p.Song]
}
func (p *Playlist) NextSong() string {
  songs := utils.GetSongs()
  p.Song++
  return songs[p.Song]
}
func (p *Playlist) OpenTrue() {
  p.Open = true 
}
func (p *Playlist) OpenFalse() {
  p.Open = false 
}
