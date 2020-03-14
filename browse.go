package vlcctrl

// Browse directory
func (instance *vlc) Browse(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/browse." + instance.Format + "?uri=" + uri)
	return
}
