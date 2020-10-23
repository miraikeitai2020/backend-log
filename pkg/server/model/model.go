package model

import(
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/db"
	"time"
)

type Log struct {
	ID            string `gorm:"primary_key"`
	UserID        int
	Date          time.Time
	Worktime      int
	Concentration []float64
	Workname	string
}

func (l *Log) Create() (e error){
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
