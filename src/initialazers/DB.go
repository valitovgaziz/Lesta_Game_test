package initialazers

import (
	"fmt"
	"lg/src/models"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PSQL models.PSQLConfig
var DB *gorm.DB

func InitDB() {
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
}