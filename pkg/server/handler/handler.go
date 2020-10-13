package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
	"log"
	"net/http"
	"github.com/google/uuid"
)

func HandleLogCreate() gin.HandlerFunc{
	return func(c *gin.Context){

		Log := model.Log{}

		logID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		Log.ID = logID.String()

		// ここに処理を書いていく
		c.JSON(200, gin.H{
			"log_id":"asdfghjkl",
		})
	}
}