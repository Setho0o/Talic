package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func FixFileNames() {
	files, err := os.ReadDir("./audio")
	if err != nil {
		log.Fatal("Error reading audio directory: ", err)
	}

	for _, e := range files {
		path := "./audio/"
		if e.IsDir() {

			ppath := "./audio/" + e.Name() + "/"

			playlistFiles, err := os.ReadDir(ppath)
			if err != nil {
				log.Fatal("Error reading "+e.Name()+" directory: ", err)
			}
			for _, pe := range playlistFiles {
				if pe.IsDir() {
					log.Fatal("Error cant have a playlist inside a playist")
				} else {
					if strings.Contains(pe.Name(), ".json") {
						continue
					}
					Rename(ppath, pe.Name())

				}
			}
		} else {
			if strings.Contains(e.Name(), ".json") {
				continue
			}
			Rename(path, e.Name())
		}
	}
}

func Rename(path, name string) {

	json, format := CheckJsonExists(name, path)
	data := JsonToMetaData(json, path)

	fmt.Println(name)
	fmt.Println(data.title+format)
	/*
	   err := os.Rename(path+name, path+fixed)

	   	if err != nil {
	   		log.Fatal("Error failed to rename: " + path + name)
	   	}
	*/
}
