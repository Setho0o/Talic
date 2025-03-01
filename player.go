package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
  "github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/ebitengine/oto/v3"
	"github.com/go-flac/go-flac/v2"
	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"
)

type SoundType int

const (
	Mp3 SoundType = iota
	Wav
  Flac
  Vorbis
  Nil
)

func DecodeAudio(t SoundType, b []byte) io.Reader {
	var reader io.Reader
	switch t {
	case Mp3:
		m, err := mp3.NewDecoder(bytes.NewReader(b))
		if err != nil {
			log.Fatal("failed decoding mp3", err)
		}
		reader = m

	case Wav:
		reader = wav.NewReader(bytes.NewReader(b))
  
  case Flac: 
    f, err := flac.ParseBytes(bytes.NewReader(b))
    if err != nil {
			log.Fatal("failed decoding flac", err)
		}
    reader = f.Frames

  case Vorbis: 
    v, err := vorbis.DecodeF32(bytes.NewReader(b))
    if err != nil {
			log.Fatal("failed decoding vorbis", err)
    } 
    reader = v 
  }
	return reader
}

func GetSoundType(path string) (SoundType, error) {
  if strings.HasSuffix(path, "mp3") {
    return Mp3, nil
  } else if strings.HasSuffix(path, "wav") {
    return Wav, nil
  } else if strings.HasSuffix(path,"flac") {
    return Flac, nil
  } else if strings.HasSuffix(path,"ogg") {
    return Vorbis, nil
  } 
  return Nil, fmt.Errorf("Invaild file must be mp3 wav or flac")
}

func Player(path string) {
	fileBytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatal("failed reading file", err)
	}

	op := &oto.NewContextOptions{}
	op.SampleRate = 44100 // Usually 44100 or 48000
	op.ChannelCount = 2 // 1 is mono sound, and 2 is stereo (most speakers are stereo)
	op.Format = oto.FormatSignedInt16LE // Format of the source. go-mp3's format is signed 16bit
	otoCtx, readyChan, err := oto.NewContext(op) // Remember that you should **not** create more than one context 
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

  fileType, err := GetSoundType(path)
	if err != nil {
		log.Fatal(err)
	}
 
	player := otoCtx.NewPlayer(DecodeAudio(fileType, fileBytes))

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	player.Play()

	// We can wait for the sound to finish playing using something like this

  t := time.Now()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
    if time.Since(t) > 10 * time.Second {
      player.Close() 
    }
	}

	// Now that the sound finished playing, we can restart from the beginning (or go to any location in the sound) using seek
	// newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
	// if err != nil{
	//     panic("player.Seek failed: " + err.Error())
	// }
	// println("Player is now at position:", newPos)
	// player.Play()

	// If you don't want the player/sound anymore simply close
	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}


