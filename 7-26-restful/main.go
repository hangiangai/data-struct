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

	pen := components.New(db, app)

	pen.Use("restful", map[string]string{
		"tname": "student",
		"mid":   "request|token",
		"mod":   "student",
	})

	pen.Use("user", map[string]string{
		"mid": "request",
	})
	app.Run(iris.Addr(":8000"))

}
