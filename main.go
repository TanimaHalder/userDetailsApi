package main

import (
	"net/http"
	"userDetailsApi/router"
	"userDetailsApi/store"
)

func main() {

	db := store.New()
	r := router.NewRouter(db)

	s := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}
	s.ListenAndServe()

}
