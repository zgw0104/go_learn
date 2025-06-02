package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// token 由三部分组成:
// 头部:（包含令牌类型如'jwt',和签名算法如'hs256')
// 载荷:(包含三类声明：1、注册声明，2、公共声明，3、私有声明)
// 签名：将编码后的Header和Payload用指定算法加密后生成：
// HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload),secret)

var mySecret = []byte("zgw0104")

const TokenExpireTime = time.Hour * 2

type MyClaim struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (atoken, rtoken string, err error) {
	claims := MyClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireTime)),
			//签发者
			Issuer: "gw",
		},
	}

	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)

	//refresh token
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)),
		Issuer:    "gw",
	}).SignedString(mySecret)

	return atoken, rtoken, err
}

func ParseToken(tokenString string) (*MyClaim, error) {
	var mc = new(MyClaim)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

//
//func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
//	if _, err = jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
//		return mySecret, nil
//	}); err != nil {
//		return
//	}
//
//	//从旧 atoken中解析出claim数据
//	var claims MyClaim
//	_, err = jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (interface{}, error) {
//		return mySecret, nil
//	})
//	v, _ := err.(*jwt.ValidationError)
//
//	if v.Errors == jwt.ValidationErrorExpired {
//		return GenerateToken(claims.UserId)
//	}
//	return
//}
