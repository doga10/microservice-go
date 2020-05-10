package status

import "github.com/gorilla/mux"

func StatusRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/", GetStatus).Methods("GET")
	return router
}
