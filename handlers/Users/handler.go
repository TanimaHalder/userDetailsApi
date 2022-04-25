package Users

import (
	"fmt"
	"net/http"
	"userDetailsApi/store"
	"userDetailsApi/types"
)

type us struct {
	db *store.DB
}

func New(db *store.DB) *us {
	return &us{db: db}
}
func (u us) NewUser(rw http.ResponseWriter, r *http.Request) {

	var user types.User
	if err := user.DecodeFromJSON(r.Body); err != nil {
		http.Error(rw, "Failed to decode", http.StatusBadRequest)
		return
	}

	if err := u.db.InsertNewUser(&user); err != nil {
		http.Error(rw, "Failed to insert", http.StatusBadRequest)
		return
	}

	fmt.Printf("%#v", u.db.GetAllUsers())
	rw.WriteHeader(http.StatusCreated)
	user.EncodeToJSON(rw)
	return

}

func (u us) GetUserByName(rw http.ResponseWriter, r *http.Request) {

	var user types.User
	if err := user.DecodeFromJSON(r.Body); err != nil {
		http.Error(rw, "Failed to decode", http.StatusBadRequest)
		return
	}

	user = u.db.GetUserByName(&user, string(user.Name))

	fmt.Printf("%#v", u.db.GetAllUsers())
	rw.WriteHeader(http.StatusCreated)
	user.EncodeToJSON(rw)
	return

}
