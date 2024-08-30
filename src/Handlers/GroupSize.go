package handlers

import (
	"encoding/json"
	"lg/src/models"
	"net/http"
)

var Size = 0

func GroupSize(w http.ResponseWriter, r *http.Request) {
	var size models.GSize
	err := json.NewDecoder(r.Body).Decode(&size)
	if err != nil {
		w.Write([]byte("can't decode json from Request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Size = size.Size
	w.WriteHeader(http.StatusOK)
}
