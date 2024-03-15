package service

import (
	"encoding/json"
	"net/http"

	"github.com/vmihailenco/msgpack/v5"
)

func ParseRequest(r *http.Request, v interface{}) error {
	if r.Header.Get("Content-Type") == "application/msgpack" {
		return msgpack.NewDecoder(r.Body).Decode(v)
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func EncodeResponse(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if r.Header.Get("Accept") == "application/msgpack" {
		w.Header().Set("Content-Type", "application/msgpack")
		return msgpack.NewEncoder(w).Encode(v)
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
