package response

import (
	"encoding/json"
	"net/http"
)

func ServeJSON(w http.ResponseWriter, data interface{}, status int) {
	res := response{
		status: status,
		Data:   data,
	}
	res.serveJSON(w)
}

type response struct {
	status int
	Data   interface{} `json:"data,omitempty"`
}

func (res *response) serveJSON(w http.ResponseWriter) {
	if res.status == 0 {
		res.status = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
