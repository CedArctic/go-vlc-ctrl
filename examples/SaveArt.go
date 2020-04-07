package main

import (
	"fmt"
	vlcctrl "github.com/CedArctic/go-vlc-ctrl"
	"io/ioutil"
)

func main() {

	// Create instance of local VLC
	myVLC, err := vlcctrl.NewVLC("127.0.0.1", 8080, "password")
	if err != nil {
		fmt.Printf("Error connecting to VLC instance:\n %s", err)
		return
	}

	// Add item to playlist
	err = myVLC.Add("file:///D:/Jose/Music/All%20Night-Lionel%20Richie.mp3")
	if err != nil {
		fmt.Printf("Error adding item to playlist:\n %s", err)
		return
	}

	// Download and save art
	var data []byte
	data, err = myVLC.Art()
	if err != nil {
		fmt.Printf("Error retrieving Art:\n %s", err)
		return
	}
	err = ioutil.WriteFile("Art.jpg", data, 0666)
	if err != nil {
		fmt.Printf("Error writing file:\n %s", err)
		return
	}
}
