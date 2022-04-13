package controller

import (
	"CopyFelix/common"
	"CopyFelix/model"
	"CopyFelix/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"xorm.io/xorm"
)

var db *xorm.Engine

func Register(context *gin.Context) {

	db = common.InitDb()
	er := db.Sync2(new(model.User))

	if er != nil {
		log.Default()
	}
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
	//
	//newUser := new(model.User)
	//newUser.Name = name
	//newUser.Password = password
	//newUser.Telephone = telephone
	hasedPassword, er2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if er2 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "加密错误",
		})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	affected, err := db.Insert(&newUser)
	if err != nil {
		fmt.Println(affected)
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

func Login(context *gin.Context) {
	//获取参数
	DB := common.InitDb()
	//telephone = ctx.PostForm("telephone")
	telephone := context.PostForm("Telephone")
	password := context.PostForm("password")
	//数据验证
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
	//判断手机号是否存在
	var user model.User
	user.Telephone = telephone
	get, err := DB.Get(&user)
	if err != nil {
		return
	}
	if !get {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	//发放token
	token := "11"
	context.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登陆成功",
	})
	//
}
