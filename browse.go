package vlcctrl

// Browse directory of provided URI
func (instance *VLC) Browse(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/browse." + instance.Format + "?uri=" + uri)
	return
}
