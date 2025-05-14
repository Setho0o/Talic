#! /usr/bin/bash
[ "$1" = "-h" -o "$1" = "--help" -o "$1" = "" ] && echo "

  format : mp3, wav, flac, vorbis; Defaults to wav.
  url: only for one song.
  
  ex: ytdlp wav https://youtube.com/... 
  

  playlist : will bring you through setup
  url: only for playlists

  ex: ytdlp playlist https://youtube.com/playlist...
  

  ffmpeg is required for ytdlp to work with the audio files.
  This is just a simple wrapper for the yt-dlp repo,
  go there for more options.

  https://github.com/yt-dlp/yt-dlp

" && exit

Default="wav" #default format
SongDir="audio"

#  -o "%(title)s.%(ext)s"\
ytdlp() {
  yt-dlp \
  -P $3 \
  --parse-metadata "description:(?s)(?P<meta_comment>.+)" \
  --parse-metadata "title:%(artist)s - %(title)s"\
  --parse-metadata "%(webpage_url)s:%(url)s" \
  --parse-metadata "%(like_count)s:%(meta_likes)s" \
  --parse-metadata "%(view_count)s:%(meta_view)s" \
  --parse-metadata "%(average_rating)s:%(rating)s" \
  --replace-in-metadata "title,uploader" "[ -]" "_" \
  --embed-metadata --merge-output-format mkv --write-info-json \
  -x --audio-format $1 $2 
}


if [ "$1" = "playlist" ]; then
  echo "choose format, defaults to" $Default
  read i
  echo $i
  if [ $i = "" ]; then
    $i = $Default
  fi
  echo "Choose playlist name"
  read j  
  mkdir audio/$j
  ytdlp $i $2 "audio/"$j
  exit  
fi
if [ "$1" = "mp3" -o "$1" = "wav" -o "$1" = "flac" -o "$1" = "vorbis" ]; then
  echo "formatting to" $1
  ytdlp $1 $2 $SongDir
else 
  echo "defaulting to" $Default
  ytdlp $Default $1 $SongDir
fi


