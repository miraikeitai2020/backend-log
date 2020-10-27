package model

import (
	"fmt"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/db"
)

type Log struct {
	LogID         string `gorm:"primary_key"`
	UserID        string
	Date          string
	WorkTime      string
	Concentration string
	LogName       string
}

func (l *Log) Create() error{
	database := db.GetDB()
	return database.Create(l).Error
}

func (l *Log) FindAllLogIDByUserID() []string {
	database := db.GetDB()
	var LogID []string
	database.Where("user_id = ?",l.UserID).Model(&Log{}).Pluck("log_id",LogID)
	fmt.Println(LogID)
	return LogID
}

func (l *Log) FindLogNameByLogID(s string) string{
	database := db.GetDB()
	var LogName string
	database.Where("log_id = ?", s).Select("log_name").Find(LogName)
	fmt.Println(LogName)
	return LogName
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
