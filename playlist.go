package vlcctrl

import "encoding/json"

// Node structure (node or leaf type) is the basic element of VLC's playlist tree representation.
// Leafs are playlist items. Nodes are playlists or folders inside playlists.
type Node struct {
	Ro       string `json:"ro"`
	Type     string `json:"type"` // node or leaf
	Name     string `json:"name"`
	ID       string `json:"id"`
	Duration int    `json:"duration,omitempty"`
	URI      string `json:"uri,omitempty"`
	Current  string `json:"current,omitempty"`
	Children []Node `json:"children,omitempty"`
}

// ParsePlaylist parses Playlist() responses to Node
func ParsePlaylist(playlistResponse string) (playlist Node, err error) {
	err = json.Unmarshal([]byte(playlistResponse), &playlist)
	if err != nil {
		return
	}
	return
}

// Returns the currently loaded playlists
func (instance *VLC) Playlist() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/playlist.json")
	return
}
