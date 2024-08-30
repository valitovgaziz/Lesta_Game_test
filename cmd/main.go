package main

import (
	"log/slog"

	intls "lg/src/initialazers"
)


func main() {

	intls.InitRouting()
	intls.InitDB()

	slog.Info("server is closed", "info", <- intls.Done)
	slog.Info("End")
}
