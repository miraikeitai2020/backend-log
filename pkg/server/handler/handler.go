package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/view"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/utils"
	"log"
	"net/http"
)

func HandleLogCreate() gin.HandlerFunc { // ログ情報追加
	return func(c *gin.Context) {
		// uuidの生成とLogにstringで保存する処理をかく
		Log:= model.Log{}
		var err error
		Log.LogID, err = utils.CreateUUIDToken()
		if err != nil {
			log.Println(err)
			view.NewErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Cannot create token",
				)
			return
		}

		// requestをバインドする
		userid := c.Request.Header.Get("x-token")
		request := view.RequestCreateLog{}
		if err := c.BindJSON(&request); err != nil {
			log.Println(err)
			view.NewErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Request is failed:" + err.Error(),
				)
			return
		}
		request.UserID = userid

		// requestを成形してmodel.Logに渡す
		request.ParseModel(&Log)

		// INSERT INTO `logs` (`log_id`,`user_id`,`date`,`work_time`,`concentration`,`log_name`) VALUES ('','','','','','')
		Log.Create()

		// responseの成形
		view.NewResponseCreateLog(c, &Log)

	}

}

func HandleLogGet() gin.HandlerFunc{ // ログ情報単体取得
	return func (c *gin.Context) {

		LogID := c.Request.Header.Get("x-token")
		if LogID == "" {
			log.Println("x-token is empty")
			view.NewErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Please enter x-token",
				)
			return
		}
		fmt.Println("log_id = " + LogID)

		Logs, err := model.FindByLogID(LogID)
		if err != nil {
			log.Println(err)
			view.NewErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Database related error",
				)
			return
		}
		view.NewResponseGetLog(c,Logs)

	}
}

func HandleLogsGet() gin.HandlerFunc{ // ログ情報一覧取得
	return func(c *gin.Context){
		// headerからuser_idを持ってくる
		Log := model.Log{}
		Log.UserID = c.Request.Header.Get("x-token")
		if Log.UserID == "" {
			log.Println("x-token is empty")
			view.NewErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Please enter x-token",
			)
			return
		}
		fmt.Println("user_id = " + Log.UserID)

		logIDs, err := Log.FindAllLogIDByUserID()
		if err != nil {
			log.Println(err)
			view.NewErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Database related error",
				)
			return
		}

		res := view.ResponseGetLogs{}
		for _, id := range logIDs {
			Logs, err:= model.FindByLogID(id)
			if err != nil {
				log.Println(err)
				view.NewErrorResponse(
					c,
					http.StatusInternalServerError,
					"Internal Server Error",
					"Database related error",
					)
				return
			}
			LWIA := view.Logs{LogID: id,LogName: Logs.LogName}
			res.Logs = append(res.Logs, LWIA)
		}
		c.JSON(200,res )
	}
}