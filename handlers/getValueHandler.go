package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Nikita213-hub/simpleInMemStorage/internal/storage"
)

type GetHandler struct {
	Path string
	Strg *storage.Storage
}

func (gh *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) != 1 {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request"))
		return
	}
	if !params.Has("key") {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request"))
		return
	}
	key := params.Get("key")
	val := gh.Strg.Get(key)
	if val == nil {
		w.WriteHeader(400)
		w.Write([]byte("No such value"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	res := struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}{
		Key:   key,
		Value: val,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Unexpected error"))
		return
	}
}
