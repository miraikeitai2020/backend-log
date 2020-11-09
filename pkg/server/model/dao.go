package model

import (
	"fmt"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/db"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/utils"
)

func (l *Log) Create() error{
	database := db.GetDB()
	log := l
	log.UserID = utils.CreateHashString(l.UserID)
	return database.Create(log).Error
}

func (l *Log) FindAllLogIDByUserID() ([]string, error) {
	database := db.GetDB()
	var LogID []string
	database.Debug().Where("user_id = ?",utils.CreateHashString(l.UserID)).Model(&Log{}).Pluck("log_id",&LogID)
	// エラーハンドリングをしたい
	if LogID == nil {
		return nil, dbError()
	}
	fmt.Println(LogID)
	return LogID, nil
}


func FindByLogID(LogID string) (*Log, error){
	database := db.GetDB()
	var log Log
	database.Where("log_id = ?", LogID).Find(&log)
	if log.LogName == "" || log.UserID== "" || log.Concentration == "" || log.WorkTime == 0 {

		return nil, dbError()
	}
	return &log, nil
}