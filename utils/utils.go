package utils
import (
  "os"
  "os/exec"

  "log"
  "strings"
)
func GetSongs() ([]string, []string)  {
  var songs []string 
  var playlists []string 
  entries, err := os.ReadDir("./audio/music")
  if err != nil {
    log.Fatal(err)
  } 
  for _, e := range entries {
    if strings.Contains(e.Name(), "."){
      songs = append(songs, e.Name())
    } else {
      playlists = append(playlists, e.Name()) 
    }
  }
  return songs, playlists
}


func Clear() {
  cmd := exec.Command("clear") 
  cmd.Stdout = os.Stdout
  cmd.Run()
}
