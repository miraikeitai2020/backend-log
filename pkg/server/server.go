package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/handler"
	"net/http"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func Init() {
	r := router()
	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}
	EnvErr := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if EnvErr != nil {
		http.ListenAndServe(":"+port, r)
	}

}

func router() *gin.Engine {
	r := gin.Default()

	r.POST("/api/log", handler.HandleLogCreate())

	r.GET("/api/log", handler.HandleLogGet())

	r.GET("/api/logs", handler.HandleLogsGet())

	return r
}
