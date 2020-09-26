package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/handler"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()


	log := r.Group("/log")
	{
		log.POST("/create", handler.HandleLogCreate())
	}
	return r
}