package handlers

import (
	"encoding/json"
	"lg/src/models"
	"net/http"
)

var MemoryType = "ram"

func GroupGamers(w http.ResponseWriter, r *http.Request) {
	var GamerP models.Gamer
	err := json.NewDecoder(r.Body).Decode(&GamerP)
	if err != nil {
		w.Write([]byte("Can't decode json from Request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if MemoryType == "ram" {
		models.Gamers = append(models.Gamers, GamerP)
	} // else {
	// TODO save to DB
	//}
	w.WriteHeader(http.StatusOK)
}
