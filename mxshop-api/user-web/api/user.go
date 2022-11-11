package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/response"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"mxshop-api/user-web/proto"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//HandleValidatorErr 表单验证错误处理返回
func HandleValidatorErr(c *gin.Context, err error) {
	fmt.Println(err.Error())
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errs.Translate(global.Trans),
	})
}

//HandleGrpcErrToHttp grpc状态码转http
func HandleGrpcErrToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
		}
	}
}

//GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	//获取前端数据
	Pn := c.DefaultQuery("pn", "0")
	PnInt, _ := strconv.Atoi(Pn)
	PSize := c.DefaultQuery("psize", "10")
	PSizeInt, _ := strconv.Atoi(PSize)

	fmt.Println(PnInt, PSizeInt)

	//将用户信息拿出
	claims, _ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Info("访问用户", currentUser.ID)

	//调用grpc接口
	rsp, err := global.UserClient.GetUserInfoList(context.WithValue(context.Background(), "ginContext", c), &proto.PageInfo{
		Pn:    uint32(PnInt),
		Psize: uint32(PSizeInt),
	})
	if err != nil {
		fmt.Println("地址", global.ServerConfig.UserSerInfo.Host, global.ServerConfig.UserSerInfo.Port)
		zap.S().Errorw("[GetUserList] 获取 【用户列表失败】")
		HandleGrpcErrToHttp(err, c)
		return
	}

	result := make([]interface{}, 0)
	for _, v := range rsp.Data {
		data := response.UserResponse{
			Id:       v.Id,
			NickName: v.NickName,
			BirthDay: response.JsonTime(time.Unix(int64(v.Birthday), 0)),
			Gender:   v.Gender,
			Mobile:   v.Mobile,
		}

		result = append(result, data)
	}
	c.JSON(http.StatusOK, result)
}

//PassWordLogin 密码登录
func PassWordLogin(c *gin.Context) {
	//表单验证, PasswordLoginForm用于存储登录数据
	PasswordLoginForm := forms.PassWordLoginForm{}
	//ShouldBind数据解析绑定,将数据放入PasswordLoginForm
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		fmt.Println(err.Error())
		HandleValidatorErr(c, err)
		return
	}

	//图片验证码验证
	if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	fmt.Println("电话号码：", PasswordLoginForm.Mobile)

	rsp, err := global.UserClient.GetUserByMobile(context.WithValue(context.Background(), "ginContext", c), &proto.MobileRequest{
		Mobile: PasswordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		//这里只是查询了用户，并没有查询密码
		passRsp, passErr := global.UserClient.CheckPassWord(context.WithValue(context.Background(), "ginContext", c), &proto.PasswordCheckInfo{
			Password:          PasswordLoginForm.Password,
			EncryptedPassword: rsp.Password,
		})
		if passErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				//登录成功，返回token
				j := middlewares.NewJWT()
				//负载内容
				Claims := models.CustomClaims{
					uint(rsp.Id),
					rsp.NickName,
					uint(rsp.Role),
					jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
						Issuer:    "ice_moss",
					},
				}
				token, err := j.CreateToken(Claims)
				if err != nil {
					zap.S().Infof("[CreateToken] 生成token失败")
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"id":        rsp.Id,
					"nick_name": rsp.NickName,
					"gender":    rsp.Gender,
					"token":     token,
					"expiresAt": (time.Now().Unix() + 60*60*24*30) * 1000,
				})

			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "登录失败",
				})
			}
		}
	}
}

//Register 用户注册
func Register(c *gin.Context) {
	//注册表单验证
	RegisterForm := forms.RegisterForm{}
	if err := c.ShouldBind(&RegisterForm); err != nil {
		HandleValidatorErr(c, err)
		return
	}

	//验证码验证
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.Redis.Host, global.ServerConfig.Redis.Port),
	})
	rsp := rdb.Get(context.Background(), RegisterForm.Mobile)
	value, err := rsp.Result()
	fmt.Println("value:", value)
	if err == redis.Nil {
		zap.S().Info("验证码错误", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return

	} else {
		if RegisterForm.Code != value {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "验证码错误",
			})

			fmt.Println("RegisterForm", RegisterForm, value)
			return
		}
	}

	userRsp, err := global.UserClient.CreateUser(context.WithValue(context.Background(), "ginContext", c), &proto.CreateUserInfo{
		NickName: fmt.Sprintf("生鲜%s", RegisterForm.Mobile),
		Password: RegisterForm.Password,
		Mobile:   RegisterForm.Mobile,
	})

	if err != nil {
		zap.S().Errorw("[Register] 创建 【用户失败】", err.Error())
		HandleValidatorErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":       "注册成功",
		"nick_name": userRsp.NickName,
	})
}

func GetStoreList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "获取商品列表成功",
	})
}
