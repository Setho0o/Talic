package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)
const (
//	folder string = "audio"
)

type MetaData struct {
  title string 
  artist string 
  url string 
  views string 
  likes string 
  date string 
  desc string
  length string 
	id string
}

func MakeSongData() map[string]MetaData {
	files, err := os.ReadDir("./audio")	
	if err != nil {
		log.Fatal("Error reading audio directory: ", err)
	}
	
	m := make(map[string]MetaData)

	for _, e := range files { 
		path := "./audio/"
		if e.IsDir() { // In Talic a directory is a playlist 
			
			ppath := "./audio/"+e.Name()+"/"

			playlistFiles, err := os.ReadDir(path)
			if err != nil {
				log.Fatal("Error reading "+e.Name()+" directory: ", err)
			}	

			for _, pe := range playlistFiles {
				if pe.IsDir() {
					log.Fatal("Error cant have a playlist inside a playist")
				} else {
					json := CheckJsonExists(pe.Name(),ppath)
					m[pe.Name()] = JsonToMetaData(json, ppath)
				}
			}

		} else {
			if strings.Contains(e.Name(),".json") {
				continue
			}
			json := CheckJsonExists(e.Name(),path)
			m[e.Name()] = JsonToMetaData(json, path)
		}
	}
  return m
}

func CheckJsonExists(file, path string) string { // returns the json file name
	// the audio files all have a corresponding json file for meta data 	
	if !strings.Contains(file,".") {
		log.Fatal("invalid file type: "+file)
	}
	n := strings.Split(file,"")
	index := strings.LastIndex(file,".")	
	fmt.Println(n[index:])

	json := strings.ReplaceAll(file, "wav","info.json")
	//json := strings.Join(n[:index],"")+".info.json"
	fmt.Println(path+json)
	_, err := os.Stat(path+json) 
	if err != nil { // if json file doesnt exist we make it 
		MakeJson(file, path)
	}

	return json
}

func MakeJson(file, path string) {
	/* 
		for now it doesnt make sense to make the json seperate from the audio files ill have to add
		music searching and move my ytdl code into golang first so for now its just an error
	*/
	log.Fatal(" json not found delete "+path+file+" and run the download script again. if the formats off by two just delete the audio file and move on idk why that shit wont work")
}

func JsonToMetaData(file, path string) MetaData {
	// for now we have to make a new payload everytime we decode a file and assert the type it would be faster to just to encode it
	// onto the struct but those are optimizations for another day
	content, err := os.ReadFile(path+file)
  if err != nil {
      log.Fatal("Error reading file: ", err)
  }

  var payload map[string]any

  err = json.Unmarshal(content, &payload)
  if err != nil {
      log.Fatal("Error during json unmashal: ", err)
  }

	return MetaData{
		title: payload["title"].(string),
		artist: payload["artist"].(string),
		url: payload["url"].(string),
		views: payload["meta_view"].(string),
		likes: payload["meta_likes"].(string),
		date: payload["upload_date"].(string),
		desc: payload["meta_comment"].(string),
		length: payload["duration_string"].(string),
		id: payload["id"].(string),
	}
}
