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

		logID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		Log := model.Log{}
		Log.ID = logID.String()

		if err := c.BindJSON(&Log); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+ err.Error())
		}

		Log.Create()

		// ここに処理を書いていく
		c.JSON(200, gin.H{
			"log_id":"asdfghjkl",// モックデータ
		})
	}
}