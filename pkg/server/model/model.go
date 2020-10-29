package model

import (
	"errors"
	"fmt"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/db"
)

type Log struct {
	LogID         string `gorm:"primary_key"`
	UserID        string
	Date          string
	WorkTime      int
	Concentration string
	LogName       string
}

func (l *Log) Create() error{
	database := db.GetDB()
	return database.Create(l).Error
}

func (l *Log) FindAllLogIDByUserID() ([]string, error) {
	database := db.GetDB()
	var LogID []string
	database.Debug().Where("user_id = ?",l.UserID).Model(&Log{}).Pluck("log_id",&LogID)
	// エラーハンドリングをしたい
	if LogID == nil {
		return nil, dbError()
	}
	fmt.Println(LogID)
	return LogID, nil
}

func (l *Log) FindByLogID(s string) (*Log, error){
	database := db.GetDB()
	var log Log
	database.Where("log_id = ?", s).Find(&log)
	if &log == nil {
		return nil, dbError()
	}
	return &log, nil
}
func dbError() error{
	return errors.New("Database related error")
}


/*
type User struct {
	ID       int `gorm:"primary_key"`
	Password string
}

type Userinfo struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Gender       int
	Age          int
	Subscription bool
}
*/
