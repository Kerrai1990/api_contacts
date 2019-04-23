package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kerrai1990/api_contacts/app"
	"github.com/kerrai1990/api_contacts/controllers"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/session", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/users/{id}/contacts", controllers.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/users/{id}/contacts", controllers.CreateContact).Methods("POST")

	err := http.ListenAndServe(":8089", router)
	if err != nil {
		fmt.Print(err)
	}
}
