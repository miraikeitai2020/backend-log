package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
	err error
)

func Init(models ...interface{}) {
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

	autoMigration(models...)

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

func autoMigration(models ...interface{}) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(models...) // ここにマイグレートしたいモデルを構造体を入れる
}
