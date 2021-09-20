package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type TestResponse struct {
	Message string
	Token   string `json:"authToken,omitempty"`
}

// Login Generate a token for a valid user
func (c *Controller) Hello(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path + " Endpoint Hit")
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if r.Header != nil {
		debug := false
		for name, values := range r.Header {
			for _, value := range values {
				if strings.ToLower(name) == "x-debug" {
					if strings.ToLower(value) == "true" {
						debug = true
					}
					break
				}
			}
		}
		for name, values := range r.Header {
			for _, value := range values {
				if !debug {
					logger.Info("%v: %v", name, value)
				} else {
					if strings.Contains(name, "X-Debug") {
						logger.Info("%v: %v", name, value)
					}
				}
			}
		}
	}
	vars := mux.Vars(r)
	key := vars["name"]
	if key == "" {
		key = "stranger"
	}
	message := fmt.Sprintf("Hello %v", key)
	response := TestResponse{
		Message: message,
	}
	if len(authHeader) == 2 {
		response.Token = authHeader[1]
	}

	json.NewEncoder(w).Encode(response)
}
