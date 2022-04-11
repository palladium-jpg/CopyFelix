package common

import (
	"fmt"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitDb() *xorm.Engine {
	driverName := "mysql"
	//host := "localhost"
	//port := "3306"
	database := "ginhello"
	username := "root"
	password := "20030531lizhe"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=true",
		username,
		password,
		database,
		charset)

	db, err := xorm.NewEngine(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	return db
}

func GetDB() *xorm.Engine {
	return engine
}
