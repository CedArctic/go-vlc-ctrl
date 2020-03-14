package vlcctrl

// Get the full list of VLM elements
func (instance *vlc) Vlm() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/vlm.xml")
	return
}

// Execute VLM Command
func (instance *vlc) VlmCmd(cmd string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/vlm_cmd.xml?command=" + cmd)
	return
}

// Get last VLM Error
func (instance *vlc) VlmCmdErr() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/vlm_cmd.xml")
	return
}
