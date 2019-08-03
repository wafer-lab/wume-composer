package core

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"wume-composer/internal/pkg/config"
	"wume-composer/internal/pkg/controllers"
	"wume-composer/internal/pkg/db"
	"wume-composer/internal/pkg/middleware"
)

func StartApp() error {
	if err := db.Open(); err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// BASE ROUTER

	router.Use(middleware.ApplyCors)
	router.Use(middleware.PanicCatcher)

	router.HandleFunc("/", controllers.IndexHandler)

	// API ROUTER

	apiRouter.Use(middleware.AuthChecker)
	apiRouter.Use(middleware.ApplyJsonContentType)

	apiRouter.HandleFunc("/", controllers.ApiIndexHandler)
	apiRouter.HandleFunc("/session", controllers.IsAuth).Methods("GET", "OPTIONS")
	apiRouter.HandleFunc("/session", controllers.SignIn).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/session", controllers.SignOut).Methods("DELETE", "OPTIONS")
	apiRouter.HandleFunc("/password", controllers.UpdatePassword).Methods("PUT", "OPTIONS")

	// STATIC
	router.PathPrefix("/static").Handler(http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./static")),
	))

	fmt.Println("Starting core at " + config.Core.Port)
	return http.ListenAndServe(":"+config.Core.Port, router)
}

func StopApp() {
	fmt.Println("Stopping core...")
	if err := db.Close(); err != nil {
		fmt.Println(err)
	}
}
