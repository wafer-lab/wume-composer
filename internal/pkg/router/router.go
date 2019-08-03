package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	c "wume-composer/internal/pkg/controllers"
)

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Group struct {
	Prefix string
	Routes []Route
	Groups []Group
}

var routes = Group{
	Routes: []Route{
		{Path: "", Method: "GET", Handler: c.ApiIndexHandler},
		{Path: "/password", Method: "PUT", Handler: c.UpdatePassword},
		// {Path: "/avatar", Method: "PUT", Handler: c.UpdateAvatar},
		// {Path: "/users", Method: "GET", Handler: c.GetUser},
	},
	Groups: []Group{
		{Prefix: "/session", Routes: []Route{
			{Path: "", Method: "GET", Handler: c.IsAuth},
			{Path: "", Method: "POST", Handler: c.SignIn},
			{Path: "", Method: "DELETE", Handler: c.SignOut},
		}},
		{Prefix: "/user", Routes: []Route{
			// {Path: "", Method: "GET", Handler: c.GetUser},
			// {Path: "/{id:[0-9]+}", Method: "GET", Handler: c.GetUser},
			// {Path: "", Method: "POST", Handler: c.CreateUser},
			// {Path: "", Method: "PUT", Handler: c.UpdateUser},
			// {Path: "", Method: "DELETE", Handler: c.RemoveUser},
		}},
	},
}

func initGroup(router *mux.Router, group Group) {
	subRouter := router.PathPrefix(group.Prefix).Subrouter()
	for _, route := range group.Routes {
		subRouter.HandleFunc(route.Path, route.Handler).Methods(strings.ToUpper(route.Method))
	}
	for _, child := range group.Groups {
		initGroup(subRouter, child)
	}
}

func InitRouter(prefix string) *mux.Router {
	routes.Prefix = prefix
	router := mux.NewRouter()
	initGroup(router, routes)
	log.Println("router has been initialized")
	return router
}
