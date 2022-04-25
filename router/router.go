package router

import (
	"net/http"
	"userDetailsApi/handlers/Users"
	"userDetailsApi/store"

	"github.com/gorilla/mux"
)

func NewRouter(db *store.DB) *mux.Router {
	r := mux.NewRouter()

	userH := Users.New(db)
	r.HandleFunc("/UserDetails", userH.NewUser).Methods(http.MethodPost)
	r.HandleFunc("/fetchUserByName", userH.GetUserByName).Methods(http.MethodPost)
	return r
}
