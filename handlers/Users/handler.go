package Users

import (
	"fmt"
	"net/http"
	"userDetailsApi/store"
	"userDetailsApi/types"
	"github.com/gin-gonic/gin"
)

type us struct {
	db *store.DB
}

func New(db *store.DB) *us {
	return &us{db: db}
}

func (u us) NewUserAdd(c *gin.Context) {

	var user types.User
	if err := user.DecodeFromJSON(c.Request.Body); err != nil {
		http.Error(c.Writer, "Failed to decode", http.StatusBadRequest)
		return
	}

	if err := u.db.InsertNewUser(&user); err != nil {
		http.Error(c.Writer, "Failed to insert", http.StatusBadRequest)
		return
	}

	fmt.Printf("%#v", u.db.GetAllUsers())
	c.Writer.WriteHeader(http.StatusCreated)
	user.EncodeToJSON(c.Writer)
	return

}

func (u us) GetUserByName(c *gin.Context) {

	var user types.User
	if err := user.DecodeFromJSON(c.Request.Body); err != nil {
		http.Error(c.Writer, "Failed to decode", http.StatusBadRequest)
		return
	}

	user = u.db.GetUserByName(&user, string(user.Name))

	fmt.Printf("%#v", u.db.GetAllUsers())
	c.Writer.WriteHeader(http.StatusCreated)
	user.EncodeToJSON(c.Writer)
	return

}
