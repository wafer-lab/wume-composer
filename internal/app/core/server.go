package core

import (
	"log"
	"net/http"

	"wume-composer/internal/pkg/db"
	"wume-composer/internal/pkg/middleware"
	"wume-composer/internal/pkg/router"
)

type Params struct {
	Port   string
	Prefix string
}

func StartApp(params Params) error {
	if err := db.Open(); err != nil {
		log.Println(err)
	}

	r := router.InitRouter(params.Prefix)

	// Middleware
	r.Use(middleware.ApplyCors)
	r.Use(middleware.PanicCatcher)
	r.Use(middleware.AuthChecker)
	r.Use(middleware.ApplyJsonContentType)

	// TODO: Move static handler to NGINX
	// Static files
	r.PathPrefix("/static").Handler(http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./static")),
	))

	log.Println("Starting core at " + params.Port)
	return http.ListenAndServe(":"+params.Port, r)
}

func StopApp() {
	log.Println("Stopping core...")
	if err := db.Close(); err != nil {
		log.Println(err)
	}
}
