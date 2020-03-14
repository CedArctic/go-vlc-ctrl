package vlcctrl

import (
	"errors"
	"strconv"
)

// Status
func (instance *vlc) Status() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format)
	return
}

// Play playlist item with given id. If id is omitted, play last active item
func (instance *vlc) Play(itemID ...int) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(itemID) > 1 {
		err = errors.New("please provide only up to one ID")
		return
	}
	urlSegment := "/requests/status." + instance.Format + "?command=pl_play"
	if len(itemID) == 1 {
		urlSegment = urlSegment + "&id=" + strconv.Itoa(itemID[0])
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Pause: If current state was 'stop', play item with given id, if no id specified, play current item.
// If no current item, play 1st item in the playlist.
func (instance *vlc) Pause(itemID ...int) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(itemID) > 1 {
		err = errors.New("please provide only up to one ID")
		return
	}
	urlSegment := "/requests/status." + instance.Format + "?command=pl_pause"
	if len(itemID) == 1 {
		urlSegment = urlSegment + "&id=" + strconv.Itoa(itemID[0])
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Stop
func (instance *vlc) Stop() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_stop")
	return
}

// Next
func (instance *vlc) Next() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_next")
	return
}

// Previous
func (instance *vlc) Previous() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_previous")
	return
}

// Empty Playlist
func (instance *vlc) EmptyPlaylist() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_empty")
	return
}

// Random Playback Toggle
func (instance *vlc) ToggleLoop() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_random")
	return
}

// Playback Looping Toggle
func (instance *vlc) ToggleRepeat() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_loop")
	return
}

// Playback Repeat Toggle
func (instance *vlc) ToggleRandom() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_repeat")
	return
}

// Playback Repeat Toggle
func (instance *vlc) ToggleFullscreen() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=fullscreen")
	return
}

// Add URI to playlist and start playback. the option field is optional, and can have the values: noaudio, novideo
// URI string is expected to be percent-encoded like URLs
func (instance *vlc) AddStart(uri string, option ...string) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(option) > 1 {
		err = errors.New("please provide only one option")
		return
	}
	urlSegment := "/requests/status." + instance.Format + "?command=in_play&input=" + uri
	if len(option) == 1 {
		if (option[0] != "noaudio") && (option[0] != "novideo") {
			err = errors.New("invalid option")
			return
		}
		urlSegment = urlSegment + "&option=" + option[0]
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Add URI to playlist. URI string is expected to be percent-encoded like URLs
func (instance *vlc) Add(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=in_enqueue&input=" + uri)
	return
}
