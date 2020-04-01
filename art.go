package vlcctrl

import (
	"errors"
	"strconv"
)

// Art fetches cover art based on a playlist item's ID. If no ID is provided, Art returns the current item's cover art.
func (instance *VLC) Art(itemID ...int) (byteArr []byte, statusCode int, err error) {

	// Check variadic arguments
	if len(itemID) > 1 {
		err = errors.New("please provide only up to one ID")
		return
	}

	// Build request URL
	urlSegment := "/art"
	if len(itemID) == 1 {
		urlSegment = urlSegment + "?item=" + strconv.Itoa(itemID[0])
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
