package vlcctrl

// Status
func (instance *vlc) Status() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format)
	return
}

// Play
func (instance *vlc) Play() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_play")
	return
}

// Pause
func (instance *vlc) Pause() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_pause")
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
