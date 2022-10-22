package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v\n", err)
		return nil
	}
	return engine
}
