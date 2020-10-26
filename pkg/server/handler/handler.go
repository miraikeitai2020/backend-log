package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/view"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HandleLogCreate() gin.HandlerFunc{
	return func(c *gin.Context){
		// uuidの生成とLogにstringで保存する処理をかく
		logID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		Log := model.Log{}
		Log.LogID = logID.String()

		// requestをバインドする
		request := view.RequestCreateLog{}
		if err := c.BindJSON(&request); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+ err.Error())
		}
		// requestで受け取ったデータをそのままブチ込める奴はここでブチ込む
		Log.UserID = request.UserID
		Log.LogName = request.LogName

		// requestで受け取ったdatetime型をstringに変換してLogにブチ込む
		Log.Date = request.Date.String()
		Log.WorkTime = request.WorkTime.String()
		// concentrationを文字列化し，結合する
		var list []string
		for i := range request.Concentration {
			list = append(list, strconv.FormatFloat(request.Concentration[i], 'f', -1, 64))
		}
		Log.Concentration = strings.Join(list, ",")

		Log.UserID = request.UserID


		Log.Create()

		response := view.ResponseCreateLog{}
		response.LogID = logID.String()
		// ここに処理を書いていく
		c.JSON(200, response)
	}
}