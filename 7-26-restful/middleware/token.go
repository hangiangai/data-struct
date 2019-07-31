package middleware

import (
	util "7-26-restful/utils"
	"time"

	"github.com/kataras/iris"
)

//token验证中间件
func CheckToken(ctx iris.Context) {
	//获取token
	token := ctx.GetHeader("token")
	var secret_key string
	row := util.Database().QueryRow("select secret_key from users where token=?", token)
	if row.Scan(&secret_key) != nil { //token不存在
		ctx.JSON(map[string]interface{}{
			"code":    "401",
			"message": "token is failure",
		})
		return
	}
	token_info := util.Get(token, secret_key)
	if val, ok := token_info["exp"].(float64); ok {
		if time.Now().Unix() < int64(val) {
			ctx.Next()
		} else {
			ctx.JSON(map[string]interface{}{
				"code":    "401",
				"message": "token is failure",
			})
		}
	}
}
