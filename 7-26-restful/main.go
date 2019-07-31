package main

import (
<<<<<<< HEAD
	"7-26-restful/components"
	util "7-26-restful/utils"
=======
	"7-26-restful/model"
	"7-26-restful/restful"
	"database/sql"
	"encoding/base64"
	"fmt"
>>>>>>> 98ce2d52778f90c26f1fdbccd5c2c9dd34e352f7

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

func main() {

	//连接数据库
	db := util.DatabaseConnection()
	app := iris.New()

<<<<<<< HEAD
	cos := components.New(db, app)
=======
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
>>>>>>> 98ce2d52778f90c26f1fdbccd5c2c9dd34e352f7

	cos.Use("restful", map[string]string{
		"tablename": "student",
	})
	cos.Use("restful", map[string]string{
		"tablename": "teacher",
	})

	app.Run(iris.Addr(":8000"))

}
