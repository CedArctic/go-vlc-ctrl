package vlcctrl

import (
	"net/url"
)

// Vlm returns the full list of VLM elements
func (instance *VLC) Vlm() (response string, err error) {
	response, err = instance.RequestMaker("/requests/vlm.xml")
	return
}

// VlmCmd executes a VLM Command and returns the response. Command is internally URL percent-encoded
func (instance *VLC) VlmCmd(cmd string) (response string, err error) {
	response, err = instance.RequestMaker("/requests/vlm_cmd.xml?command=" + url.QueryEscape(cmd))
	return
}

// VlmCmdErr returns the last VLM Error
func (instance *VLC) VlmCmdErr() (response string, err error) {
	response, err = instance.RequestMaker("/requests/vlm_cmd.xml")
	return
}
