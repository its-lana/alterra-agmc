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

func WrapResponse(code int, status string, response interface{}) *WebResponse {
	newResponse := new(WebResponse)
	newResponse.Code = code
	newResponse.Status = status
	newResponse.Data = response
	return newResponse
}

func (resp *WebResponse) WriteToResponseBody(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	return json.NewEncoder(w).Encode(resp)
}
