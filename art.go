package vlcctrl

import (
	"errors"
	"strconv"
)

// Fetch cover art: Use itemID = 0 for currently playing
func (instance *vlc) Art(itemID int) (byteArr []byte, statusCode int, err error) {

	// Build request URL
	urlSegment := "/art"
	if itemID != 0 {
		urlSegment = urlSegment + "?item=" + strconv.Itoa(itemID)
	}

	// Make request
	var response string
	response, byteArr, statusCode, err = instance.RequestMaker(urlSegment)

	// Error Handling
	if err != nil {
		return
	}
	if response == "Error" {
		err = errors.New("no cover art available for item")
	}
	return
}
