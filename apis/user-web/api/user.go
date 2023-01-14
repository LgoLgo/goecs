package api

import (
	"apis/user-web/forms"
	"apis/user-web/global"
	"apis/user-web/global/response"
	"apis/user-web/models"
	"apis/user-web/validator"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	middlewares "apis/user-web/middleware"

	proto "apis/user-web/proto/gen"
)

func HandleGRPCErrorToHTTP(err error, c *app.RequestContext) {
	// Convert gRPC code to HTTP status code
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, utils.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg:": "Internal error",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, utils.H{
					"msg": "Parameter error",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg": "User service unavailable",
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
	claims, _ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("User: %d", currentUser.ID)

	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.UserSrvClient.GetUserList(ctx, &proto.PageInfo{
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
	// form validation
	validator.ValidateMobile()
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.BindAndValidate(&passwordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		c.JSON(http.StatusBadRequest, utils.H{
			"captcha": "wrong code",
		})
		return
	}

	// Login logic
	if rsp, err := global.UserSrvClient.GetUserByMobile(ctx, &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "User doesn't exist",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "Login failed",
				})
			}
			return
		}
	} else {
		// Have searched user
		if passRsp, err := global.UserSrvClient.CheckPassWord(ctx, &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "Login failed",
			})
		} else {
			if passRsp.Success {
				// Generate token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               // Signature valid time
						ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30 days expired
						Issuer:    "ecs",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, utils.H{
						"msg": "Generate token failed",
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
					"msg": "Login failed",
				})
			}
		}
	}
}

func Register(ctx context.Context, c *app.RequestContext) {
	validator.ValidateMobile()
	// User register
	registerForm := forms.RegisterForm{}
	if err := c.BindAndValidate(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})

	value, err := rdb.Get(ctx, registerForm.Mobile).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"code": "wrong code",
		})
		return
	} else {
		if value != registerForm.Code {
			fmt.Printf("value: %s\ncode: %s", value, registerForm.Code)
			c.JSON(http.StatusBadRequest, utils.H{
				"code": "wrong code",
			})
			return
		}
	}

	user, err := global.UserSrvClient.CreateUser(ctx, &proto.CreateUserInfo{
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
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    "ecs",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{
			"msg": "Generate token failed",
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
