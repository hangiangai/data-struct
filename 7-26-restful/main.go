package main

import (
	"7-26-restful/model"
	"7-26-restful/restful"
	"database/sql"
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

	r.Init(db, app)

	r.Register(model.Student{}, "student")
	r.Register(model.Teacher{}, "teacher")

	app.Run(iris.Addr(":8000"))
}
