package router

import (
	"userDetailsApi/handlers/Users"
	"userDetailsApi/store"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *store.DB) *gin.Engine {
	r := gin.Default()

	userH := Users.New(db)
	r.POST("/UserDetails", userH.NewUserAdd)

	r.POST("/fetchUserByName", userH.GetUserByName)

	return r
}
