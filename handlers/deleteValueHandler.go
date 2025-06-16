package handlers

import (
	"net/http"

	"github.com/Nikita213-hub/simpleInMemStorage/internal/storage"
)

type DeleteHandler struct {
	Strg *storage.Storage
}

func (dh *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	err := dh.Strg.Delete(key)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Record was deleted succefully"))
}
