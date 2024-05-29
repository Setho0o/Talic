package tui

import (
  "fmt"
  "github.com/Setho0o/Talic/audio"
)

func Debug(playlist audio.Playlist) {

  fmt.Println(playlist.GetSongs(),playlist.CurrentPlaylist(), "\n", playlist.Song, "\t", playlist.Time)
}
