package handlers

import (
	"net/http"

	"github.com/Nikita213-hub/simpleInMemStorage/internal/storage"
)

type DeleteHandler struct {
	Strg *storage.Storage
}

func (dh *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	if len(key) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request"))
		return
	}
	err := dh.Strg.Delete(key)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Record was deleted succefully"))
}
