package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

var Gamers = make([]Gamer, 10, 15)
var MemoryType = "ram"

func main() {
	slog.Info("Init routing")
	r := chi.NewRouter()
	r.Post("/users", GroupGames)
	r.Post("/memory", Memory)
	fmt.Println("Main.go")
}

func GroupGames(w http.ResponseWriter, r *http.Request) {
	var GamerP Gamer
	err := json.NewDecoder(r.Body).Decode(&GamerP)
	if err != nil {
		w.Write([]byte("Can't decode json from Request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if MemoryType == "ram" {
		Gamers = append(Gamers, GamerP)
	} else {
		TODO save to DB
	}
	w.WriteHeader(http.StatusOK)
}

func Memory(w http.ResponseWriter, r *http.Request) {
	var memory MemType
	err := json.NewDecoder(r.Body).Decode(&memory)
	if err != nil {
		w.Write([]byte("can't decode json from Request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	MemoryType = memory.Type
	w.WriteHeader(http.StatusOK)
}

type Gamer struct {
	Name    string `json:"name"`
	Skill   float32 `json:"skill"`
	Latency float32 `json:"latency"`
}

type MemType struct {
	Type string `json:"type"`
}

type PSQLConfig struct {
	PGUSER string
	PGPASSWORD string
	PGDATABASE string
	PGPORT string
	PGHOST string
}