package controller

import (
	"CopyFelix/common"
	"CopyFelix/model"
	"CopyFelix/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xorm.io/xorm"
)

func Register(context *gin.Context) {

	db := common.GetDB()

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
		name = util.RandomString(10)
	}
	if isTelephoneExist(db, telephone) {
		context.JSON(http.StatusUnprocessableEntity, "用户已经存在")
		return
	}
	//log.Panicln(name, telephone, password)

	newUser := new(model.User)
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

}

func isTelephoneExist(db *xorm.Engine, telephone string) bool {
	var user model.User
	user.Telephone = telephone
	//_, err := db.Where("telephone = ?", telephone).Exist(user)
	get, err := db.Get(&user)
	if err != nil {
		return get
	}
	return get
}
