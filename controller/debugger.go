package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DebugResponse struct {
	Token         string   `json:"token"`
	Headers       []string `json:"headers"`
	Ip            string   `json:"ip"`
	Host          string   `json:"host"`
	Url           string   `json:"url"`
	RemoteAddress string   `json:"remoteAddress"`
	RequestUri    string   `json:"requestUri"`
	Proto         string   `json:"proto"`
}

// Login Generate a token for a valid user
func (c *Controller) Debug(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path + " Endpoint Hit")
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

	debugResponse := DebugResponse{}

	debugResponse.Ip = r.RemoteAddr
	debugResponse.Host = r.Host
	debugResponse.RemoteAddress = r.RemoteAddr
	debugResponse.RequestUri = r.RequestURI
	debugResponse.Url = r.URL.Host
	debugResponse.Token = authHeader[0]
	debugResponse.Proto = r.Proto

	if r.Header != nil {
		for name, values := range r.Header {
			for _, value := range values {
				header := fmt.Sprintf("%v: %v", name, value)
				debugResponse.Headers = append(debugResponse.Headers, header)
				logger.Info("%v: %v", name, value)
			}
		}
	}

	logger.Info("IP: %v", debugResponse.Ip)
	logger.Info("Host: %v", debugResponse.Host)
	logger.Info("RemoteAddress: %v", debugResponse.RemoteAddress)

	json.NewEncoder(w).Encode(debugResponse)
}
