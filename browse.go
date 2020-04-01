package vlcctrl

import (
	"encoding/json"
)

// Element represents a single item in the browse directory
type Element struct {
	ItemType         string `json:"type"`
	Path             string `json:"path"`
	AccessTime       uint   `json:"access_time"`
	UID              uint   `json:"uid"`
	CreationTime     uint   `json:"creation_time"`
	GID              uint   `json:"gid"`
	ModificationTime uint   `json:"modification_time"`
	Mode             uint   `json:"mode"`
	URI              string `json:"uri"`
	Size             uint   `json:"size"`
}

// Browse is an intermediate structure used for parsing Element structs
type Browse struct {
	Elements []Element `json:"element"`
}

// Browse directory of provided URI
func (instance *VLC) Browse(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/browse.json?uri=" + uri)
	return
}

// ParseBrowse parses Browse() responses to []Element
func ParseBrowse(browseResponse string) (files []Element, err error) {
	var browse Browse
	err = json.Unmarshal([]byte(browseResponse), &browse)
	if err != nil {
		return
	}
	files = browse.Elements
	return files, nil
}
