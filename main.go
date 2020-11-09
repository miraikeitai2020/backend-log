package main

import (
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/db"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
)

func main() {
	db.Init(model.Log{})
	//
	/*if err := db.Init(model.Log{}); err != nil {
		log.Panic("DB connection error")
	}*/
	server.Init()
	db.Close()
}
