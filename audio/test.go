package audio

import (
	"log"
	"os"
  "time"  
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
)

func Player(song string) {
	// Open the audio file
	file, err := os.Open(song)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the audio file
	streamer, format, err := flac.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize the speaker
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	// Play the audio
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Wait for the audio to finish playing
	<-done
}
