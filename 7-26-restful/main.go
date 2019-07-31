package main

import (
	"7-26-restful/components"
	util "7-26-restful/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

func main() {

	//连接数据库
	db := util.DatabaseConnection()
	app := iris.New()

	cos := components.New(db, app)

	cos.Use("restful", map[string]string{
		"tablename": "student",
	})
	cos.Use("restful", map[string]string{
		"tablename": "teacher",
	})

	app.Run(iris.Addr(":8000"))

}
