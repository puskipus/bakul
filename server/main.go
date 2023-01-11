package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hanifhahn/bakul/model"
	"github.com/hanifhahn/bakul/routes"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	routes.ProdukRoutes(r)
	model.ConnectDB()

	fmt.Println("Hello world")

	r.Run(":8080")
}
