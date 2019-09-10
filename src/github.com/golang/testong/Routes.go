package main

import "net/http"

//Route is entity of routing
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        bool
}

//Routes is array of Route
type Routes []Route

/*
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"ProductCreate",
		"POST",
		"/products",
		ProductCreate,
	},
	Route{
		"ProductUpdate",
		"PUT",
		"/products",
		ProductUpdate,
	},
}
*/
