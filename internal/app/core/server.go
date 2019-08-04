package core

import (
	"net/http"

	"wume-composer/internal/pkg/db"
	"wume-composer/internal/pkg/logger"
	"wume-composer/internal/pkg/middleware"
	"wume-composer/internal/pkg/router"
)

type Params struct {
	Port   string
	Prefix string
}

func StartApp(params Params) error {
	if err := db.Open(); err != nil {
		logger.Error(err.Error())
	}

	r := router.InitRouter(params.Prefix)

	// Middleware
	r.Use(middleware.PanicCatcher)
	r.Use(middleware.Logger)
	r.Use(middleware.ApplyJsonContentType)
	r.Use(middleware.ApplyCors)
	r.Use(middleware.AuthChecker)

	// TODO: Move static handler to NGINX
	// Static files
	r.PathPrefix("/static").Handler(http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./static")),
	))

	logger.Info("Starting core at " + params.Port)
	return http.ListenAndServe(":"+params.Port, r)
}

func StopApp() {
	logger.Info("Stopping core...")
	if err := db.Close(); err != nil {
		logger.Error(err.Error())
	}
}
