package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	DBMS := "mysql"
	USER := "root"
	PASS := "okirahirem"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "merihariko"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	autoMigration()

	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}

func Close(){
	if err := db.Close(); err != nil{
		panic(err)
	}
}

func autoMigration() {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate() // ここにマイグレートしたいモデルを構造体を入れる
}
