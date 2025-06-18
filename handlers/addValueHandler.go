package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Nikita213-hub/simpleInMemStorage/internal/storage"
)

type AddHandler struct {
	Path string
	Strg *storage.Storage
}

func (ah *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect content type"))
		return
	}
	key := r.PathValue("key")
	if len(key) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect query params"))
		return
	}
	bodyDecoced := struct {
		Value interface{} `json:"value"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&bodyDecoced)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Unexpected error"))
		return
	}
	err = ah.Strg.Put(key, bodyDecoced.Value)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Unexpected error"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Value was added successfully"))
}
