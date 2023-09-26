package main

import (
	"go_pro/common"

	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	ginServer := gin.Default()
	CollectRoutes(ginServer)
	ginServer.Run(":8080")
}
