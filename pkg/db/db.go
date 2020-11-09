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
	/*
	DBMS := "mysql"
	USER := "root"
	PASS := "okirahirem"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "merihariko"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	*/
	dbms := "mysql"
	connect := "b76627ce158cf5:13b3cecb@tcp(us-cdbr-east-02.cleardb.com:3306)/heroku_331a3da834c9b7e?parseTime=true"
	db, err = gorm.Open(dbms, connect)
	if err != nil {
		panic(err.Error())
	}

	autoMigration(models...)

	db.LogMode(true)
}

/*func Init(models ...interface{}) {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("MYSQL_HOST")
	PORT := os.Getenv("MYSQL_PORT")
	DATABASE := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DATABASE

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	autoMigration(models...)

	db.LogMode(true)
}*/

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
	/*for _, m := range models{
		if !db.HasTable(m) {
			db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(m)
		}
	}*/
}
