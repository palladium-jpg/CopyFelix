package common

import (
	"CopyFelix/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "CopyFile",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	return token, claims, err
}

//jwt json web token 为了在网络应用环境间传递声明的一种基于JSON的开放标准（(RFC 7519).
//该token被设计为紧凑且安全的，特别适用于分布式站点的单点登录（SSO）场景。
//JWT的声明一般被用来在身份提供者和服务提供者间传递被认证的用户身份信息，
//以便于从资源服务器获取资源，也可以增加一些额外的其它业务逻辑所必须的声明信息，该token也可直接被用于认证，也可被加密。
//https://www.jianshu.com/p/576dbf44b2a 详细内容可参考

/*


 */
//
