package components

import (
	util "7-26-restful/utils"
	"crypto/md5"
	"database/sql"
	"time"

	"github.com/kataras/iris"
)

var (
	sql_statement = []string{
		"select username from users where username=?",
		"insert into users (username,password,oid,secret_key,created_at) values (?,?,?,?,?)",
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
	u.database = cos.Db //将值赋给全局对象
	cos.App.AllowMethods(iris.MethodOptions)
	router := cos.App.Party("/user/", cos.Mid...) //创建路由组
	router.Post("/login", u.Login)                //用户登录
	router.Post("/register", u.Register)          //用户注册
}

//注册
func (u User) Register(ctx iris.Context) {

	reg_data := make(map[string]string)
	ctx.ReadJSON(&reg_data)
	//1.用户提交数据错误
	if _, ok := Filter(reg_data, []string{"username", "password"}); !ok {
		ctx.JSON(Response{
			Code:    404,
			Message: "the request failed",
			Data:    "problems parsing json",
		})
		return
	}

	var username string
	row := u.database.QueryRow(sql_statement[0], reg_data["username"])
	//2.用户已存在
	if row.Scan(&username) == nil {
		ctx.JSON(Response{
			Code:    404,
			Message: "the user already exists",
			Data:    reg_data["username"],
		})
		return
	}
	//对象对应的id
	oid := md5.Sum([]byte(reg_data["username"]))
	//用户对token对应的key
	secret_key := oid[3:15]
	//插入数据 插入用户名,密码,用户id号,初始秘钥
	result, err := u.database.Exec(sql_statement[1], reg_data["username"], reg_data["password"], oid, secret_key, time.Now())
	util.ErrProcess(err, 58)
	rowsAffected, err := result.RowsAffected()
	util.ErrProcess(err, 80)

	//3.插入用户失败
	if rowsAffected <= 0 {
		ctx.JSON(Response{
			Code:    404,
			Message: "registration failed",
			Data:    reg_data["username"],
		})
		return
	}
	//4.注册成功
	ctx.JSON(Response{
		Code:    200,
		Message: "registered successfully",
		Data:    oid,
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
	if _, ok := Filter(login_data, []string{"username", "password"}); !ok {
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

func Filter(v map[string]string, key []string) (string, bool) {
	if len(key) > 0 && v != nil {
		for _, val := range key {
			if val_, ok := v[val]; !ok || len(val_) == 0 {
				return val + " empty", false
			}
		}
	}
	return "ok", true
}
