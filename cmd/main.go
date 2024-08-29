package main

import (
	"fmt"
	hdl "lg/src/Handlers"
	"lg/src/models"
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

var Gamers = make([]models.Gamer, 10, 15)
var MemoryType = "ram"
var Size = 0
var Done = make(chan bool)
var PSQL models.PSQLConfig
var DB *gorm.DB

func main() {
	// up server and chi routing
	slog.Info("Init routing")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/users", hdl.GroupGamers)   // resive json Gamer{name: string, skill: num, latency: num}
	r.Post("/memory", hdl.Memory)       // resive json memType{Type: string}
	r.Post("/groupSize", hdl.GroupSize) // resive json GSize{size: number}
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
	if err != nil {
		slog.Error("failed to open db", "error", err)
		os.Exit(2)
	}
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
