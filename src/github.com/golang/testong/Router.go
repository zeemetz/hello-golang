package main

import "github.com/gin-gonic/gin"

func init() {
	admin := GetAdmin()
	admin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	admin.GET("/testong/:testong/:message", func(c *gin.Context) {
		c.String(200, c.Param("testong")+c.Param("message"))
	})
}

//Gorilla Mux
//NewRouter create new router useed by main.go
// func NewRouter() *mux.Router {
// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {
// 		var handler http.Handler
//
// 		handler = route.HandlerFunc
// 		handler = Logger(handler, route.Name)
//
// 		router.
// 			Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(handler)
// 	}
//
// 	return router
// }

//NewEngine is
// func NewEngine() *gin.Engine {
// 	e, a := GetEngine()
// 	//engine
// 	e.PUT("/products", ProductUpdateGin)
// 	//basic auth access
// 	admin := a.Group("/admin")
// 	admin.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "pong"})
// 	})
// 	return e
// }
