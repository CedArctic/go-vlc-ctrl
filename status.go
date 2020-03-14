package vlcctrl

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
