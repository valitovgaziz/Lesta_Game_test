package handlers

import (
	"encoding/json"
	"lg/src/models"
	"net/http"
)

func Memory(w http.ResponseWriter, r *http.Request) {
	var memory models.MemType
	err := json.NewDecoder(r.Body).Decode(&memory)
	if err != nil {
		w.Write([]byte("can't decode json from Request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	MemoryType = memory.Type
	w.WriteHeader(http.StatusOK)
}
