package oauth

import "github.com/gorilla/mux"

func OAuthRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/oauth", nil).Methods("GET")
	return router
}
