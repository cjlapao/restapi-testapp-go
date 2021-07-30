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
	logger.Info("Test Endpoint Hit")
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
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
