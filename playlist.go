package vlcctrl

// Status
func (instance *VLC) Playlist() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/playlist." + instance.Format)
	return
}
