package audio

import (
  "os"
  "log"
  "strings"
)

type Playlist struct {
  Playlist int
  Song int
  Pause bool
  Time int 
}

func (p *Playlist) GetPlaylists() []string  {
  var playlists []string 
  entries, err := os.ReadDir("./audio/music/")
  if err != nil {
    log.Fatal(err)
  } 
  for _, e := range entries {
    if strings.Contains(e.Name(), "."){
    } else {
      playlists = append(playlists, e.Name()) 
    }
  }
  return playlists
}

func (p *Playlist) GetSongs() []string {
  var songs []string 
  entries, err := os.ReadDir(p.CurrentPlaylist())
  if err != nil {
    log.Fatal(err)
  } 
  for _, e := range entries {
    songs = append(songs, e.Name())
  }
  return songs
}
func (p *Playlist) CurrentPlaylist() string {
  playlists := p.GetPlaylists()
  return "audio/music/" + playlists[p.Playlist]
}
func (p *Playlist) CurrentSong() string {
  songs := p.GetSongs()
  return songs[p.Song]
}

func (p *Playlist) NextPlaylist() {
  if p.Playlist == len(p.GetPlaylists()) - 1  { 
    p.Playlist = 0
  } else {
    p.Playlist++
  }
}
func (p *Playlist) PreviousPlaylist() {
  if p.Playlist <= 0  { 
    p.Playlist = len(p.GetPlaylists()) - 1
  } else {
    p.Playlist--
  }
}
func (p *Playlist) NextSong() {
  s := p.GetSongs()
  if p.Song == len(s) - 1  { 
    p.Song = 0
  } else {
    p.Song++
  }
}
func (p *Playlist) PreviousSong() {
  s := p.GetSongs()
  if p.Song <= 0  { 
    p.Song = len(s) - 1
  } else {
    p.Song--
  }
}
func (p *Playlist) SetTime(int) {
  p.Time = p.Time / 1000000000
}



