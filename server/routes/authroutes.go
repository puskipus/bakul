package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanifhahn/bakul/controllers/authcontroller"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", authcontroller.Login)
	r.POST("/register", authcontroller.Register)
	r.GET("/logout", authcontroller.Logout)
}
