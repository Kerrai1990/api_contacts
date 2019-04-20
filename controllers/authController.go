package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kerrai1990/api_contacts/models"
	u "github.com/kerrai1990/api_contacts/utils"
)

// CreateUser -
var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))
	}

	response := account.Create()
	u.Respond(w, response)
}

// Authenticate -
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))
	}

	response := models.Login(account.Email, account.Password)
	u.Respond(w, response)
}
