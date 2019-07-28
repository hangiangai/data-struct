package main

import (
	"7-26-restful/model"
	"7-26-restful/restful"
	"database/sql"
	"encoding/base64"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "hangiangai"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}

	app := iris.New()

	r := restful.New()

	routes := app.Party("/user")

	routes.Get("/", func(ctx iris.Context) {
		ctx.WriteString("1234567")
	})

	routes.Post("/register", func(ctx iris.Context) {

		register_info := make(map[string]string)
		ctx.ReadJSON(&register_info)

		fmt.Println(register_info)
		var username string //存用户名

		//处理账号
		if val, ok := register_info["username"]; ok && username != "" {
			row := db.QueryRow("select username from users where username=", val)
			row.Scan(&username)
			if len(username) != 0 { //该用户存在,返回
				return
			}
		}

		fmt.Println(username)

		//处理密码
		if password, ok := register_info["password"]; ok && password != "" && username == "" {
			base64 := base64.StdEncoding.EncodeToString([]byte(password))

			fmt.Println(base64)
			register_info["secret_key"] = base64[7:14]
			register_info["objectId"] = base64[0:6]
			sql := restful.CreateInsertSql("users", register_info)
			_, err := db.Exec(sql)
			restful.ErrProcess(err, 32)
			ctx.WriteString("yes")
		}

	})

	fmt.Println("123")

	r.Init(db, app)

	r.Register(model.Student{}, "student")
	r.Register(model.Teacher{}, "teacher")

	app.Run(iris.Addr(":8000"))
}
