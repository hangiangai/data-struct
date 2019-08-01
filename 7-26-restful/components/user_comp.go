package components

import (
	util "7-26-restful/utils"
	"database/sql"
	"encoding/base64"

	"github.com/kataras/iris"
)

var (
	sql_statement = []string{
		"select username from users where username=?",
		"insert into users (username,password,objectId,secret_key) values (?,?,?,?)",
		"select objectId,username,password from users where username = ?",
		"update users set secret_key=?,token=? where username=?",
	}
)

type Response struct {
	Code    int64
	Message string
	Data    interface{}
	Token   string
}

//参数接受对象
type User struct {
	database *sql.DB
}

//组件
func (u User) Mount(cos Components) {
	//将值赋给全局对象
	u.database = cos.Db
	cos.App.AllowMethods(iris.MethodOptions)
	//创建路由组
	router := cos.App.Party("/user/", cos.Mid...)

	{
		//用户登录
		router.Post("/login", u.Login)
		//用户注册
		router.Post("/regiser", u.Register)
	}
}

//注册
func (u User) Register(ctx iris.Context) {

	register_data := make(map[string]string)
	ctx.ReadJSON(&register_data)
	var username string
	//1.用户提交数据错误
	if _, ok := CheckInfo(register_data, []string{"username", "password"}); !ok {
		ctx.JSON(Response{
			Code:    404,
			Message: "the request failed",
			Data:    "problems parsing json",
		})
		return
	}

	row := u.database.QueryRow(sql_statement[0], register_data["username"])
	//2.用户已存在
	//QueryRow()方法查询到返回数据,查询不到返回err
	//当查询到数据error为nil
	if row.Scan(&username) == nil {
		ctx.JSON(Response{
			Code:    404,
			Message: "the user already exists",
			Data:    register_data["username"],
		})
		return
	}

	random := base64.StdEncoding.EncodeToString([]byte(register_data["username"]))
	//插入数据 插入用户名,密码,用户id号,初始秘钥
	result, err := u.database.Exec(sql_statement[1], register_data["username"], register_data["password"], random[0:6], random[10:22])
	util.ErrProcess(err, 58)
	rowsAffected, err := result.RowsAffected()
	util.ErrProcess(err, 80)

	//3.插入用户失败
	if rowsAffected <= 0 {
		ctx.JSON(Response{
			Code:    404,
			Message: "registration failed",
			Data:    register_data["username"],
		})
		return
	}

	//4.注册成功
	ctx.JSON(Response{
		Code:    200,
		Message: "registered successfully",
		Data:    random[0:6],
	})
}

//登录
func (u User) Login(ctx iris.Context) {

	login_data := make(map[string]string)
	util.ErrProcess(ctx.ReadJSON(&login_data), 165)
	var username string
	var password string
	var objectId string
	//1.数据错误
	if _, ok := CheckInfo(login_data, []string{"username", "password"}); !ok {
		ctx.JSON(Response{
			Code:    404,
			Message: "the user already exists",
			Data:    objectId,
		})
		return
	}

	//2.账户不存在
	row := u.database.QueryRow(sql_statement[2], login_data["username"])
	util.ErrProcess(row.Scan(&objectId, &username, &password), 130)
	if len(username) <= 0 || len(password) <= 0 {
		ctx.JSON(Response{
			Code:    404,
			Message: "the user already exists",
			Data:    objectId,
		})
		return
	}
	//3.密码不正确
	if login_data["password"] != password {
		ctx.JSON(Response{
			Code:    404,
			Message: "the user already exists",
			Data:    objectId,
		})
		return
	}

	//4.用户验证成功
	aud := "browser"
	if ctx.IsMobile() {
		aud = "client"
	}
	//获取token 和关键字key (objectId:用户objectId aud:访问端 过期时间)
	token, secret_key := util.CreateToken(objectId, aud, 60*30)
	//将秘钥和token存入数据库 (秘钥,token,用户名)
	result, err := u.database.Exec(sql_statement[3], secret_key, token, username)
	util.ErrProcess(err, 208)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected > 0 {
		ctx.JSON(Response{
			Code:    200,
			Message: "user verification successful",
			Data:    objectId,
			Token:   token,
		})
	}
}

//检测所给数据中是否包含指定字段数据
func CheckInfo(v map[string]string, key []string) (string, bool) {
	if len(key) > 0 {
		for _, val := range key {
			if val_, ok := v[val]; !ok || len(val_) == 0 {
				return val + " empty", false
			}
		}
	}
	return "ok", true
}
