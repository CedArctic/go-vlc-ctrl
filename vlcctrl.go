package vlcctrl

// ===== Imports =====

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// ===== Structures =====

// Structure for an http interface enabled VLC instance
type vlc struct {
	IP       string
	Port     int
	Password string
	BaseURL  string
}

// ===== Functions =====

// Create a new VLC instance to control
func NewVLC(IP string, Port int, Password string) vlc {

	// Form instance Base URL
	var BaseURL strings.Builder
	BaseURL.WriteString("http://")
	BaseURL.WriteString(IP)
	BaseURL.WriteString(":")
	BaseURL.WriteString(strconv.Itoa(Port))
	BaseURL.WriteString("/requests/")

	// Create and return instance struct
	return vlc{IP, Port, Password, BaseURL.String()}
}

// Function used to make requests to VLC using a urlSegment
func (instance *vlc) RequestMaker(urlSegment string) (response string, statusCode int, err error) {
	// Make a GET Request
	client := &http.Client{}
	request, reqErr := http.NewRequest("GET", instance.BaseURL+urlSegment, nil)
	request.SetBasicAuth("", instance.Password)
	reqResponse, resErr := client.Do(request)

	// Get byte response and http status code
	byteArr, readErr := ioutil.ReadAll(reqResponse.Body)

	// Error handling
	if reqErr != nil {
		err = errors.New(fmt.Sprintf("HTTP request error: %s\n", reqErr))
		return
	}

	if resErr != nil {
		err = errors.New(fmt.Sprintf("HTTP response error: %s\n", reqErr))
		return
	}

	if readErr != nil {
		err = errors.New(fmt.Sprintf("Error reading response: %s\n", reqErr))
		return
	}

	// Write response and http status code
	response = string(byteArr)
	statusCode = reqResponse.StatusCode

	return
}

// Play
func (instance *vlc) Play() (response string, statusCode int, err error) {
	response, statusCode, err = instance.RequestMaker("status.json?command=pl_play")
	return
}

// Pause
func (instance *vlc) Pause() (response string, statusCode int, err error) {
	response, statusCode, err = instance.RequestMaker("status.json?command=pl_pause")
	return
}