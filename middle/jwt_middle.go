package middle

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my-take-out/common"
	"my-take-out/common/e"
	"my-take-out/common/enum"
	"my-take-out/common/utils"
	"my-take-out/global"
	"net/http"
)

func VerifyJWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS

		// 这样获取token，要自定义请求头
		token := c.Request.Header.Get(global.Config.Jwt.Admin.Name)

		// 这样获取token，在请求参数里找到auth就行
		//token := strings.TrimSpace(c.GetHeader("Authorization")) // 先去除首尾空格
		//token = strings.TrimPrefix(token, "Bearer")
		//token = strings.TrimSpace(token) // 再次清理可能的多余空格

		headers, _ := json.Marshal(c.Request.Header)
		fmt.Printf("Headers: %s\n", headers)

		if token == "" {
			c.JSON(401, gin.H{"error": "Token format invalid"})
			c.Abort()
			return
		}
		log.Println("Received token:", token) // 打印接收到的令牌

		payload, err := utils.ParseToken(token, global.Config.Jwt.Admin.Secret)
		if err != nil {
			log.Println("Token parsing error:", err) // 打印解析错误
			code = e.UNKNOW_IDENTITY                 // 实际返回的是401未鉴权
			c.JSON(http.StatusUnauthorized, common.Result{Code: code})
			c.Abort()
			return
		}

		c.Set(enum.CurrentId, payload.UserId)
		c.Set(enum.CurrentName, payload.GrantScope)

		c.Next()
	}
}
