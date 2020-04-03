package vlcctrl

import (
	"net/url"
)

// Get the full list of VLM elements
func (instance *VLC) Vlm() (response string, err error) {
	response, err = instance.RequestMaker("/requests/vlm.xml")
	return
}

// Execute VLM Command. Command is internally URL percent-encoded
func (instance *VLC) VlmCmd(cmd string) (response string, err error) {
	response, err = instance.RequestMaker("/requests/vlm_cmd.xml?command=" + url.QueryEscape(cmd))
	return
}

// Get last VLM Error
func (instance *VLC) VlmCmdErr() (response string, err error) {
	response, err = instance.RequestMaker("/requests/vlm_cmd.xml")
	return
}
