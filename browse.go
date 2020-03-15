package vlcctrl

// Browse directory based on its URI. URI is expected to be percent encoded
func (instance *VLC) Browse(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/browse." + instance.Format + "?uri=" + uri)
	return
}
