package model

import (
	"time"
)

type Log struct {
	ID            int `gorm:"primary_key"`
	UserID        int
	Date          time.Time
	Worktime      time.Time
	Concentration byte
}

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
