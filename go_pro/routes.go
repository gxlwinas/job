package main

import (
	"go_pro/controller"
	"go_pro/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/captcha", controller.Captcha)
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/publish", controller.JobPublic)
	r.GET("/jobinfo", controller.JobInfo)
	r.PUT("/jobupdata/:id", controller.JobUpdata)
	r.DELETE("jobdelete/:id", controller.JobDelete)

	return r
}
