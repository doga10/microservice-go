package status

import (
	"microservice-go/framework/handler/response"
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	response.RenderJSON(w, 200, "APP Service Available! GoLang: v1.0.0", true)
}
