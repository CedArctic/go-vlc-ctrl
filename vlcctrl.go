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
type VLC struct {
	IP       string
	Port     int
	Password string
	BaseURL  string
	Format   string // json or xml
}

// ===== Functions =====

// Create a new VLC instance to control
func NewVLC(IP string, Port int, Password string, Format string) VLC {

	// Form instance Base URL
	var BaseURL strings.Builder
	BaseURL.WriteString("http://")
	BaseURL.WriteString(IP)
	BaseURL.WriteString(":")
	BaseURL.WriteString(strconv.Itoa(Port))

	// Create and return instance struct
	return VLC{IP, Port, Password, BaseURL.String(), Format}
}

// Function used to make requests to VLC using a urlSegment
func (instance *VLC) RequestMaker(urlSegment string) (response string, byteArr []byte, statusCode int, err error) {

	// Form a GET Request
	client := &http.Client{}
	request, reqErr := http.NewRequest("GET", instance.BaseURL+urlSegment, nil)
	if reqErr != nil {
		err = errors.New(fmt.Sprintf("http request error: %s\n", reqErr))
		return
	}

	// Make a GET request
	request.SetBasicAuth("", instance.Password)
	reqResponse, resErr := client.Do(request)
	if resErr != nil {
		err = errors.New(fmt.Sprintf("http response error: %s\n", reqErr))
		return
	}
	defer reqResponse.Body.Close()

	// Get byte response and http status code
	var readErr error
	byteArr, readErr = ioutil.ReadAll(reqResponse.Body)
	if readErr != nil {
		err = errors.New(fmt.Sprintf("error reading response: %s\n", reqErr))
		return
	}

	// Write response and http status code
	response = string(byteArr)
	statusCode = reqResponse.StatusCode

	return
}
