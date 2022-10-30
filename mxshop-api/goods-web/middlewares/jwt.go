package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWTAuth 验证token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					c.JSON(http.StatusUnauthorized, map[string]string{
						"msg": "授权已过期",
					})
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusUnauthorized, "未登陆")
			c.Abort()
			return
		}

		//只要验证通过，需要证明你是谁,有没有权限
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		fmt.Println("token认证成功")
		c.Next()
	}
}

//JWT 实现JWT服务
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

//NewJWT 实例化对象
func NewJWT() *JWT {
	return &JWT{
		//写入签名
		[]byte(global.ServerConfig.JWTInfo.SigningKey),
	}
}

//CreateToken 创建一个token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	//生成token的前两段
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用签名，生成token
	return token.SignedString(j.SigningKey)
}

//ParseToken 解析token包含的信息
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	//将token字符串传入,根据CustomClaims 结构体定义的相关属性要求进行校验
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

//RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
