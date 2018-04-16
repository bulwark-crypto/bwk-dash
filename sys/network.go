package sys

import (
	"encoding/json"
	"net/http"
)

// API is the api for the ip address.
// There is no currently known way to reliably get
// the public ip address for the node so we look outside.
const API string = "https://api.ipify.org?format=json"

// IPResponse holds the ip address from ipify.org.
type IPResponse struct {
	IP string
}

// GetIP will return the ip address from: https://www.ipify.org/
func GetIP() (ip string, err error) {
	var res *http.Response
	if res, err = http.Get(API); err != nil {
		return
	}

	// We could just get the string but for the sake of
	// consistency we go json.
	resIP := new(IPResponse)
	if err = json.NewDecoder(res.Body).Decode(resIP); err != nil {
		return
	}

	ip = resIP.IP
	return
}
