package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"cross-sell/sellapi"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{

	Route{
		Name:"GetMachineSKUsByType",
		Method:"GET",
		Pattern:"/machines/{machine_type}/skus",
		HandlerFunc:sellapi.GetMachineSKUsByType,
	},

	Route{
		Name:"GetPodSKUsByPodType",
		Method:"GET",
		Pattern:"/pods/{pod_type}/skus",
		HandlerFunc:sellapi.GetPodSKUsByPodType,
	},

	Route{
		Name:"GetCrossSellSKUsByQueryParam",
		Method:"GET",
		Pattern:"/crosssell/skus",
		HandlerFunc:sellapi.GetCrossSellSKUsByQueryParam,
	},
}

// NewRouter function configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}