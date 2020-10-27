package model

import (
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

func (l *Log) Create() (e error) {
	database := db.GetDB()
	return database.Create(l).Error
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
