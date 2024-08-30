package initialazers

import (
	hdl "lg/src/Handlers"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Done = make(chan bool)

func InitRouting() {

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

}
