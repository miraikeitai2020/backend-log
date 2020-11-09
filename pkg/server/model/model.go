package model

import (
	"errors"
)

type Log struct {
	LogID         string `gorm:"primary_key"`
	UserID        string
	Date          string
	WorkTime      int
	Concentration string
	LogName       string
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
