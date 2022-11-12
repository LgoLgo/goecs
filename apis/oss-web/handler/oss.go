package handler

import (
	"apis/oss-web/global"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"net/url"
	"strings"

	"apis/oss-web/utils"
)

func Token(ctx context.Context, c *app.RequestContext) {
	response := utils.GetPolicyToken()
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, response)
}

func Request(c context.Context, ctx *app.RequestContext) {
	fmt.Println("\nHandle Post Request ... ")
	// Get PublicKey bytes
	bytePublicKey, err := utils.GetPublicKey(ctx)
	if err != nil {
		utils.ResponseFailed(ctx)
		return
	}

	// Get Authorization bytes : decode from Base64String
	byteAuthorization, err := utils.GetAuthorization(ctx)
	if err != nil {
		utils.ResponseFailed(ctx)
		return
	}

	// Get MD5 bytes from Newly Constructed Authorization String.
	byteMD5, bodyStr, err := utils.GetMD5FromNewAuthString(ctx)
	if err != nil {
		utils.ResponseFailed(ctx)
		return
	}

	decodeUrl, err := url.QueryUnescape(bodyStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeUrl)
	params := make(map[string]string)
	datas := strings.Split(decodeUrl, "&")
	for _, v := range datas {
		sdatas := strings.Split(v, "=")
		fmt.Println(v)
		params[sdatas[0]] = sdatas[1]
	}
	fileName := params["filename"]
	fileUrl := fmt.Sprintf("%s/%s", global.ServerConfig.OssInfo.Host, fileName)

	// verifySignature and response to client
	if utils.VerifySignature(bytePublicKey, byteMD5, byteAuthorization) {
		// do something you want according to callback_body ...
		ctx.JSON(http.StatusOK, utils2.H{
			"url": fileUrl,
		})
		//utils.ResponseSuccess(ctx)  // response OK : 200
	} else {
		utils.ResponseFailed(ctx) // response FAILED : 400
	}
}
