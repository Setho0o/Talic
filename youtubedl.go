package main

import (
	"fmt"
	"log"
"github.com/lithdew/nicehttp"
	"github.com/lithdew/youtube"
)

func DownloadMp3() {
  results, err := youtube.Search("doomsday", 0)
  if err != nil {
    log.Fatal("failed to find song")
  }
	fmt.Printf("Got %d search result(s).\n\n", results.Hits)

	if len(results.Items) == 0 {
		log.Fatal("got zero search results")
	}
	// Get the first search result and print out its details.

	details := results.Items[0]
	fmt.Printf(
		"ID: %q\n\nTitle: %q\nAuthor: %q\nDuration: %q\n\nView Count: %q\nLikes: %d\nDislikes: %d\n\n",
		details.ID,
		details.Title,
		details.Author,
		details.Duration,
		details.Views,
		details.Likes,
		details.Dislikes,
	)

	// Instantiate a player for the first search result.

	player, err := youtube.Load(details.ID)
	if err != nil {
    log.Fatal("cant create player")
  }
	// Fetch audio-only direct link.

	stream, ok := player.SourceFormats().AudioOnly().BestAudio()
	if !ok {
		log.Fatal("no audio-only stream available")
	}

  audioOnlyFilename := "audio." + stream.FileExtension()

	audioOnlyURL, err := player.ResolveURL(stream)
  if err != nil {
    log.Fatal("")
  }
  nicehttp.DownloadFile(audioOnlyFilename, audioOnlyURL)
}
