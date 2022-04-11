package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"net/http"
	"time"
	"xorm.io/xorm"
)

type User struct {
	//gorm.Model
	Name      string `xorm:"type:varchar(20);not null"`
	Telephone string `xorm:"varchar(110);not null;primary key"`
	Password  string `xorm:"varchar(110);not null"`
}

var engine *xorm.Engine

func main() {
	//fmt.Println("Hello World!!!")
	//db := InitDb()
	//defer db.Close()
	db := InitDb()
	er := db.Sync2(new(User))
	if er != nil {
		log.Default()
	}
	r := gin.Default()
	r.POST("/ping", func(context *gin.Context) {

		name := context.PostForm("Name")

		telephone := context.PostForm("Telephone")

		password := context.PostForm("password")

		if len(telephone) != 11 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":      422,
				"msg":       "手机号码必须为11位",
				"telephone": telephone,
			})
			return
		}

		if len(password) < 6 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不能少于六位",
			})
			return
		}

		if len(name) == 0 {
			name = RandomString(10)
		}
		if isTelephoneExist(db, telephone) {
			context.JSON(http.StatusUnprocessableEntity, "用户已经存在")
			return
		}
		//log.Panicln(name, telephone, password)

		newUser := new(User)
		newUser.Name = name
		newUser.Password = password
		newUser.Telephone = telephone

		affeced, err := db.Insert(newUser)
		if err != nil {
			fmt.Println(affeced)
		}

		context.JSON(200, gin.H{
			"msg": "注册成功",
		})

	})
	panic(r.Run())

}
func isTelephoneExist(db *xorm.Engine, telephone string) bool {
	var user User
	user.Telephone = telephone
	//_, err := db.Where("telephone = ?", telephone).Exist(user)
	get, err := db.Get(&user)
	if err != nil {
		return get
	}
	return get
}

func RandomString(n int) string {
	var letters = []byte("huaedfuiwqehdoquiwhjdqoiuhdou")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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
