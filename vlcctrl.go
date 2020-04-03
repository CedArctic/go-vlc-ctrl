package vlcctrl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// VLC struct represents an http interface enabled VLC instance. Build using NewVLC()
type VLC struct {
	IP       string
	Port     int
	Password string
	BaseURL  string // combination of IP and Port
}

// NewVLC builds and returns a VLC struct using the IP, Port and Password of the VLC instance
func NewVLC(ip string, port int, password string) (VLC, error) {

	// Form instance Base URL
	var BaseURL strings.Builder
	BaseURL.WriteString("http://")
	BaseURL.WriteString(ip)
	BaseURL.WriteString(":")
	BaseURL.WriteString(strconv.Itoa(port))

	// Create and return instance struct
	return VLC{ip, port, password, BaseURL.String()}, nil
}

// RequestMaker make requests to VLC using a urlSegment provided by other functions
func (instance *VLC) RequestMaker(urlSegment string) (response string, err error) {

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
	defer func() {
		err = reqResponse.Body.Close()
	}()

	// Check HTTP status code and errors
	statusCode := reqResponse.StatusCode
	if !((statusCode >= 200) && (statusCode <= 299)) {
		err = errors.New(fmt.Sprintf("http error code: %s\n", statusCode))
		return
	}

	// Get byte response and http status code
	byteArr, readErr := ioutil.ReadAll(reqResponse.Body)
	if readErr != nil {
		err = errors.New(fmt.Sprintf("error reading response: %s\n", readErr))
		return
	}

	// Write response
	response = string(byteArr)

	return
}
