package main

import (
	dbModel "goworkerspace/api/model"
	"goworkerspace/api/rest"

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

//student模型
type Student struct {
	Id     string `json:id`
	Name   string `json:name`
	Age    string `json:age`
	Gender string `json:gender`
}

func main() {

	app := iris.New()

	//添加RESTful Api 方法
	restful := rest.New(app)
	restful.Add(rest.Register("student", dbModel.Init(&Student{})))

	app.Run(iris.Addr("10.197.27.47:8080"))

}
