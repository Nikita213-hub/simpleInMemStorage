package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Nikita213-hub/simpleInMemStorage/internal/storage"
)

type GetAllHandler struct {
	Strg *storage.Storage
}

func (gah *GetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	values := gah.Strg.GetAll()
	res := make([]struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}, 0, len(values))
	for k, v := range values {
		res = append(res, struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
		}{
			Key:   k,
			Value: v,
		})
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Unexpected error"))
		return
	}
}
