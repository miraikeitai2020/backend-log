package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/view"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/utils"
	"log"
	"net/http"
)

func HandleLogCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// uuidの生成とLogにstringで保存する処理をかく
		Log:= model.Log{}
		var err error
		Log.LogID, err = utils.CreateUUIDToken()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// requestをバインドする
		userid := c.Request.Header.Get("x-token")
		request := view.RequestCreateLog{}
		if err := c.BindJSON(&request); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		request.UserID = userid

		// requestを成形してmodel.Logに渡す
		request.ParseModel(&Log)

		// INSERT INTO `logs` (`log_id`,`user_id`,`date`,`work_time`,`concentration`,`log_name`) VALUES ('','','','','','')
		Log.Create()

		// responseの成形
		response := view.NewResponseCreateLog(&Log)

		// ここに処理を書いていく
		c.JSON(200, response)
	}
}

func HandleLogsGet() gin.HandlerFunc{
	return func(c *gin.Context){
		// headerからuser_idを持ってくる
		Log := model.Log{}
		Log.UserID = c.Request.Header.Get("x-token")
		if Log.UserID == "" {
			err := errors.New("Please enter x-token")
			log.Println(err)
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}
		fmt.Println("user_id = " + Log.UserID)

		logIDs, err := Log.FindAllLogIDByUserID()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		res := view.ResponseGetLogs{}
		for _, id := range logIDs {
			Logs, err:= Log.FindByLogID()
			if err != nil {
				log.Println(err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			LWIA := view.Logs{LogID: id,LogName: Logs.LogName}
			res.Logs = append(res.Logs, LWIA)
		}
		c.JSON(200,res )
	}
}