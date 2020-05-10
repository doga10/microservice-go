package users

import "github.com/gorilla/mux"

func UserRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/users", nil).Methods("GET")
	return router
}
