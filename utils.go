package main
import (
  "os"
  "log"
)
func GetSongs() []string {
  var songs []string 
  entries, err := os.ReadDir("./audio/songs")
  if err != nil {
    log.Fatal(err)
  } 
  for _, e := range entries {
    songs = append(songs, e.Name())
  }
  return songs 
}
