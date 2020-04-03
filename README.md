<div align="center">
    <img src="https://github.com/CedArctic/go-vlc-ctrl/blob/master/img/logo.png" width="140px" alt="logo"/>
</div>


# go-vlc-ctrl : Control VLC with Go
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-%2300ADD8?logo=go)](https://pkg.go.dev/github.com/CedArctic/go-vlc-ctrl)
[![godoc reference](https://img.shields.io/badge/godoc-reference-blue?logo=)](https://godoc.org/github.com/CedArctic/go-vlc-ctrl)
[![go report card](https://goreportcard.com/badge/github.com/CedArctic/go-vlc-ctrl)](https://goreportcard.com/report/github.com/CedArctic/go-vlc-ctrl)
[![GitHub license](https://img.shields.io/github/license/CedArctic/go-vlc-ctrl)](https://github.com/CedArctic/go-vlc-ctrl/blob/master/LICENSE)
[![release](https://img.shields.io/github/v/release/CedArctic/go-vlc-ctrl?color=orange)](https://github.com/CedArctic/go-vlc-ctrl/releases/latest)

A simple yet powerful module that allows you to control VLC instances over the VLC Web API. 

## Installation
Make sure you have Go installed and just open up a terminal window and run:
```bash
go get github.com/CedArctic/go-vlc-ctrl
```

## Example
To run this example, enable the VLC Web Interface (View > Add Interface > Web) and make sure it has a password set to
"password" (Preferences > All > Main interfaces > Lua > Lua HTTP > Password). 

Alternatively you can launch VLC with a Web Interface from the command prompt / terminal 
(use macosx instead of qt if on macOS):
```bash
vlc --intf http --extraintf qt --http-password password
```
```go
package main

import (
	"github.com/CedArctic/go-vlc-ctrl"
	"time"
)

func main(){
	// Declare a local VLC instance on port 8080 with password "password"
	myVLC, _ := vlcctrl.NewVLC("127.0.0.1", 8080, "password")

	// Add items to playlist. Note URIs are URL percent-encoded
	myVLC.Add("file:///C:/Users/Jose/Music/Back%%20In%%20Black.mp3")
	myVLC.Add("https://www.youtube.com/watch?v=dQw4w9WgXcQ")

	// Play first item and wait for 10 seconds
	myVLC.Play()
	time.Sleep(10 * time.Second)
	
	// Skip to next item, toggle full screen and pause after 30s
	myVLC.Next()
	myVLC.ToggleFullscreen()
	time.Sleep(30 * time.Second)
	myVLC.Pause()
}
```


## Documentation
You can find documentation of all functions on [GoDoc](https://godoc.org/github.com/CedArctic/go-vlc-ctrl) 
or [Go.dev](https://pkg.go.dev/github.com/CedArctic/go-vlc-ctrl).

The module fully covers the VLC Web API as documented 
[here](https://github.com/videolan/vlc/blob/master/share/lua/http/requests/README.txt).


## Contributing
Contributions to the project in any way are welcome

## Resources
- https://wiki.videolan.org/Interfaces/
- https://wiki.videolan.org/VLC_HTTP_requests/
- https://wiki.videolan.org/Documentation:Modules/http_intf/
- https://github.com/videolan/vlc/blob/master/share/lua/http/requests/README.txt

## License
Copyright (c) 2020 CedArctic. This project is licensed under the [MIT license](LICENSE).