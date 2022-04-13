package common

import (
	"fmt"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

func InitDb() *xorm.Engine {
	driverName := viper.GetString("datasource.driverName")
	//host := "localhost"
	//port := "3306"
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=true",
		username,
		password,
		database,
		charset)

	//if len(driverName) == 0 {
	//	fmt.Println("212")
	//} else {
	//	fmt.Println(12)
	//}
	db, err := xorm.NewEngine(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	return db
}
