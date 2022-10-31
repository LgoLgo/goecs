package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"apis/user-web/forms"
	"apis/user-web/global"
	"apis/user-web/global/response"
	middlewares "apis/user-web/middleware"
	"apis/user-web/models"
	proto "apis/user-web/proto/gen"
	"apis/user-web/validator"
)

// TODO：暂未实现选择类型登录（1选择账号密码，2选择手机验证码）

func HandleGRPCErrorToHTTP(err error, c *app.RequestContext) {
	// 将 gRPC 的 code 转换成 HTTP 的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, utils.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, utils.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *app.RequestContext, err error) {
	c.JSON(http.StatusOK, utils.H{
		"msg": err.Error(),
	})
	return
}

func GetUserList(ctx context.Context, c *app.RequestContext) {
	// 从注册中心获取用户服务的信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}

	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}

	if userSrvHost == "" {
		c.JSON(http.StatusBadRequest, utils.H{
			"captcha": "用户服务不可达",
		})
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] connected error",
			"msg", err.Error(),
		)
	}
	claims, _ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("User: %d", currentUser.ID)
	// 调用接口
	userSrvClient := proto.NewUserClient(userConn)

	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := userSrvClient.GetUserList(ctx, &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] query user list error.")
		HandleGRPCErrorToHTTP(err, c)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	c.JSON(http.StatusOK, result)
}

func PassWordLogin(ctx context.Context, c *app.RequestContext) {
	// 表单验证
	validator.ValidateMobile() // 手机号自定义表单验证设置
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.BindAndValidate(&passwordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		c.JSON(http.StatusBadRequest, utils.H{
			"captcha": "验证码错误",
		})
		return
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] connected error",
			"msg", err.Error(),
		)
	}
	// 调用接口
	userSrvClient := proto.NewUserClient(userConn)

	// 登录的逻辑
	if rsp, err := userSrvClient.GetUserByMobile(ctx, &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		// 仅查询到用户
		if passRsp, err := userSrvClient.CheckPassWord(ctx, &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				// 生成token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               // 签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
						Issuer:    "L2ncE",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, utils.H{
						"msg": "生成token失败",
					})
					return
				}
				c.JSON(http.StatusOK, utils.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}
}

func Register(ctx context.Context, c *app.RequestContext) {
	validator.ValidateMobile() // 手机号自定义表单验证设置
	//用户注册
	registerForm := forms.RegisterForm{}
	if err := c.BindAndValidate(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	//验证码
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	value, err := rdb.Get(ctx, registerForm.Mobile).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"code": "验证码错误",
		})
		return
	} else {
		if value != registerForm.Code {
			fmt.Printf("value: %s\ncode: %s", value, registerForm.Code)
			c.JSON(http.StatusBadRequest, utils.H{
				"code": "验证码错误",
			})
			return
		}
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] connected error",
			"msg", err.Error(),
		)
	}
	// 调用接口
	userSrvClient := proto.NewUserClient(userConn)

	user, err := userSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})
	if err != nil {
		zap.S().Errorf("[Register] create user error: %s", err.Error())
		HandleGRPCErrorToHTTP(err, c)
		return
	}

	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "L2ncE",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}
