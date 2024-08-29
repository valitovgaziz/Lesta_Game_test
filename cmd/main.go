package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Gamers = make([]Gamer, 10, 15)
var MemoryType = "ram"
var Size = 0
var Done = make(chan bool)
var PSQL PSQLConfig
var DB *gorm.DB

func main() {
	// up server and chi routing
	slog.Info("Init routing")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/users", GroupGamers)   // resive json Gamer{name: string, skill: num, latency: num}
	r.Post("/memory", Memory)       // resive json memType{Type: string}
	r.Post("/groupSize", GroupSize) // resive json GSize{size: number}
	go func() {
		defer close(Done)
		err := http.ListenAndServe(":3000", r)
		if err != nil {
			slog.Error("Can't statrt server ", "error", err)
		}
	}()
	slog.Info("server started on 3000")

	// init DB connection
	err := cleanenv.ReadConfig("config.env", &PSQL)
	if err != nil {
		slog.Error("Can't read config", "error", err)
		os.Exit(2)
	}

	slog.Info(PSQL.PGDATABASE)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Yekaterinburg",
		PSQL.PGHOST,
		PSQL.PGUSER,
		PSQL.PGPASSWORD,
		PSQL.PGDATABASE,
		PSQL.PGPORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	DB, err := db.DB()
	if err != nil {
		slog.Error("failed to get database", "error", err)
		os.Exit(2)
	}
	err = DB.Ping()
	if err != nil {
		slog.Error("failed to ping database", "error", err)
		os.Exit(2)
	}
	slog.Info("connected to database")

	slog.Info("server is closed", "info", <-Done)
	slog.Info("End")
}

func GroupGamers(w http.ResponseWriter, r *http.Request) {
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
	PGUSER     string `env:"PGUSER"`
	PGPASSWORD string `env:"PGPASSWORD`
	PGDATABASE string `env:"PGDATABASE"`
	PGPORT     string `env:"PGPORT"`
	PGHOST     string `env:"PGHOST"`
}
