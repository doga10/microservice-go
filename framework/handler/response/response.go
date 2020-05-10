package response

import (
	"encoding/json"
	"net/http"
)

type handlerJSON struct {
	Code   int         `jwt:"code"`
	Data   interface{} `jwt:"data"`
	Status bool        `jwt:"status"`
}

func RenderJSON(w http.ResponseWriter, code int, data interface{}, status bool) {
	w.Header().Set("Content-Type", "application/jwt")
	handler := handlerJSON{Code: code, Status: status, Data: data}
	render, err := json.Marshal(handler)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(render)
	return
}
