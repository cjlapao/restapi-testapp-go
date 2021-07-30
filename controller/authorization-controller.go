package controller

import (
	"encoding/json"
	"io/ioutil"
	"ittech24/rest/apidemo/entities"
	"ittech24/rest/apidemo/log"
	"ittech24/rest/apidemo/security"
	"net/http"
)

// Login Generate a token for a valid user
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	logger := log.Get()
	logger.Debug("Login Endpoint Hit")

	reqBody, _ := ioutil.ReadAll(r.Body)
	loginRequest := entities.LoginRequest{}
	json.Unmarshal(reqBody, &loginRequest)

	user := entities.User{
		Email:    "admin@localhost",
		Password: "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
	}
	if len(user.Email) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	password := security.SHA256Encode(loginRequest.Password)

	if password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, expires := security.GenerateUserToken(user.Email)
	response := entities.LoginResponse{
		AccessToken: string(token),
		Expiring:    expires,
	}

	json.NewEncoder(w).Encode(response)
}

// Validate Validate a token for a valid user
func (c *Controller) Validate(w http.ResponseWriter, r *http.Request) {
	token, valid := security.GetAuthorizationToken(r.Header)

	if !valid {
		response := entities.LoginErrorResponse{
			Code:    "404",
			Error:   "Token Not Found",
			Message: "The JWT token was not found or the header was malformed, please check your authorization header",
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	if !security.ValidateToken(token) {
		response := entities.LoginErrorResponse{
			Code:    "401",
			Error:   "Invalid Token",
			Message: "The JWT token is not valid",
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := entities.LoginResponse{
		AccessToken: token,
	}

	json.NewEncoder(w).Encode(response)
}