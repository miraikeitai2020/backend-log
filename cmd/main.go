package main

import (
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/db/"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server"
)

func main(){
	db.Init()
	server.Init()
	db.Close()
}
