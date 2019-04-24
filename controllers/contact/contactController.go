package contact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kerrai1990/api_contacts/models"
	u "github.com/kerrai1990/api_contacts/utils"
)

// CreateContact -
func CreateContact(w http.ResponseWriter, r *http.Request) {

	// Get User from Context
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))
	}

	contact.UserID = user
	response := contact.Create()
	u.Respond(w, response)
}

// GetContactsFor -
func GetContactsFor(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "Id not provided"))
		return
	}

	fmt.Println(params)
	data := models.GetUserContacts(uint(id))
	response := u.Message(true, "Success")
	response["data"] = data
	u.Respond(w, response)

}
