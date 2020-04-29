/*
What this example does: https://www.youtube.com/watch?v=Bl9uBYA7c1U
Scenario: You're in the office with 100 computers all on the same subnet with IPs ranging from 192.168.1.1
to 192.168.1.100. They all have VLC with the web interface enabled on port 8080 and the password set to "password". You
run this and suddenly the office is filled with 100 Gandalfs nodding to the tune of Epic Sax Guy.
*/

package main

import (
	"fmt"
	vlcctrl "github.com/CedArctic/go-vlc-ctrl"
	"time"
)

func main() {
	// Create an array of 100 VLC instances
	var gandalfs [100]vlcctrl.VLC

	// Add the video to the playlist of each PC
	for i := 1; i < 100; i++ {
		gandalfs[i], _ = vlcctrl.NewVLC(fmt.Sprintf("192.1.1.%d", i), 8080, "password")
		gandalfs[i].Add("https://youtu.be/G1IbRujko-A")
	}

	// Hit play
	for i := 1; i < 100; i++ {
		gandalfs[i].Play()
	}

	// Give time for the video to load and toggle fullscreen
	time.Sleep(10 * time.Second)
	for i := 1; i < 100; i++ {
		gandalfs[i].ToggleFullscreen()
	}
}
