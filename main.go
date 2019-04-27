package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kerrai1990/api_contacts/app"
	"github.com/kerrai1990/api_contacts/controllers/auth"
	"github.com/kerrai1990/api_contacts/controllers/contact"
	"github.com/kerrai1990/api_contacts/models"

)
 
func main() {

	fmt.Println("MAIN.GO MAIN()")
	models.MakeDB()
  
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/users", auth.CreateUser).Methods("POST")
	router.HandleFunc("/api/session", auth.Authenticate).Methods("POST")
	router.HandleFunc("/api/users/{id}/contacts", contact.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/users/{id}/contacts", contact.CreateContact).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Print(err)
	}
}
