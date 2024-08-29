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
var Size = 0

func main() {
	slog.Info("Init routing")
	r := chi.NewRouter()
	r.Post("/users", GroupGames)    // resive json Gamer{name: string, skill: num, latency: num}
	r.Post("/memory", Memory)       // resive json memType{Type: string}
	r.Post("/groupSize", GroupSize) // resive json GSize{size: number}
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
		// TODO save to DB
	}
	w.WriteHeader(http.StatusOK)
}

func GroupSize(w http.ResponseWriter, r *http.Request) {
	var size GSize
	err := json.NewDecoder(r.Body).Decode(&size)
	if err != nil {
		w.Write([]byte("can't decode json from Request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Size = size.size
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

type GSize struct {
	size int
}

type Gamer struct {
	Name    string  `json:"name"`
	Skill   float32 `json:"skill"`
	Latency float32 `json:"latency"`
}

type MemType struct {
	Type string `json:"type"`
}

type PSQLConfig struct {
	PGUSER     string
	PGPASSWORD string
	PGDATABASE string
	PGPORT     string
	PGHOST     string
}
