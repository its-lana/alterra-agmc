package models

import (
	"encoding/json"
	"net/http"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func (resp *WebResponse) WriteToResponseBody(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	return json.NewEncoder(w).Encode(resp)
}
