package restful

import (
	"encoding/base64"
	"fmt"

	"7-26-restful/util"

	"github.com/kataras/iris"
)

func UserRouter(app *iris.Application) {

	user_router := app.Party("/user", setRequestMiddleware)
	{
		//用户登录
		user_router.Post("/login", UserLogin)
		//用户注册
		user_router.Post("/register", UserReister)
		//token验证
		user_router.Post("/token", UserToken)
	}
}

//用户注册
func UserReister(ctx iris.Context) {
	//用于读取用户信息
	register_info := make(map[string]string)
	ctx.ReadJSON(&register_info)
	var username string

	//用户名存在且不为空
	if val, ok := register_info["username"]; ok && len(val) > 0 {
		//查询该用户是否存在
		row := database.QueryRow("select username from users where username=?", val)
		row.Scan(&username)
		//用户存在
		if len(username) != 0 {
			ctx.JSON(Response{
				Code:    "404",
				Message: "the user already exists",
				Data:    nil,
			})
			return
		}
		//用户不存在
		if password, ok := register_info["password"]; ok && len(password) > 0 {

			base64 := base64.StdEncoding.EncodeToString([]byte(password))

			fmt.Println(base64)
			register_info["secret_key"] = base64[7:14]     //用户秘钥
			register_info["objectId"] = base64[0:6]        //用户objectiId
			sql := CreateInsertSql("users", register_info) //生成sql语句
			result, err := database.Exec(sql)
			ErrProcess(err, 33)
			row_num, err := result.RowsAffected()
			ErrProcess(err, 35)

			if row_num != 0 {
				ctx.JSON(Response{
					Code:    "200",
					Message: "registered successfully",
					Data:    nil,
				})
			} else {
				ctx.JSON(Response{
					Code:    "404",
					Message: "registered successfully",
					Data:    "nil",
				})
			}
		}
	} else {
		ctx.JSON(Response{
			Code:    "404",
			Message: "data error",
			Data:    nil,
		})
	}
}

//用户登录
func UserLogin(ctx iris.Context) {

	login_info := make(map[string]string)
	ctx.ReadJSON(&login_info)
	var username string
	var password string

	if _, ok := CheckInfo(login_info, []string{"username", "password"}); ok {
		//查询用户名和密码
		row := database.QueryRow("select username,password from users where username = ?", login_info["username"])
		row.Scan(&username, &password)
		if len(username) > 0 { //存在该用户
			if login_info["password"] == password {

				token, secret_key := util.CreateToken(username)

				result, err := database.Exec("update users set secret_key=?,token=? where username=?", secret_key, token, username)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(result.RowsAffected())

				ctx.JSON(Response{
					Code:    "200",
					Message: "yes",
					Data:    token,
				})

			} else {
				ctx.JSON(Response{
					Code:    "404",
					Message: "password or username error",
					Data:    login_info["username"],
				})
			}
		} else { //用户不存在
			ctx.JSON(Response{
				Code:    "404",
				Message: "users don't exist",
				Data:    login_info["username"],
			})
		}
	} else {
		ctx.JSON(Response{
			Code:    "404",
			Message: "data error",
			Data:    nil,
		})
	}

}

func CheckInfo(v map[string]string, key []string) (string, bool) {
	if len(key) > 0 {
		for _, val := range key {
			if val_, ok := v[val]; !ok || len(val_) == 0 {
				return val + " empty", false
			}
		}
	}
	return "successful", true
}

//token验证
func UserToken(ctx iris.Context) {
	token := ctx.GetHeader("token")
	var username string
	var secret_key string
	//根据token查找用户和key
	row := database.QueryRow("select username,secret_key from users where token = ?", token)
	row.Scan(&username, &secret_key)
	if len(secret_key) > 0 {
		h, p := util.Get(token, secret_key)
		fmt.Println(h)
		fmt.Println(p)
		ctx.JSON(h)
	}
}
