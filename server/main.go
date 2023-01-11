package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hanifhahn/bakul/middlewares"
	"github.com/hanifhahn/bakul/model"
	"github.com/hanifhahn/bakul/routes"
)

func main() {
	// make gin router and use logger for logging terminal
	r := gin.New()
	r.Use(gin.Logger())

	// make auth router
	routes.AuthRoutes(r)

	// make produk router
	v1 := r.Group("/api")
	v1.Use(middlewares.JWTMiddleware())
	routes.ProdukRoutes(v1)

	// connect database
	model.ConnectDB()

	// run server on localhost:8080
	r.Run(":8080")
}
