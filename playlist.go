package vlcctrl

// Status
func (instance *vlc) Playlist() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/playlist." + instance.Format)
	return
}
