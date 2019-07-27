package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
)

func newApp() *iris.Application {
	app := iris.New()
	//授权配置
	authConfig := basicauth.Config{
		Users:   map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
		Realm:   "Authorization Required",
		Expires: time.Duration(30) * time.Minute,
	}
	authentication := basicauth.New(authConfig)

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/admin")
	})

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", h)
		// http://localhost:8080/admin/profile
		needAuth.Get("/profile", h)
		// http://localhost:8080/admin/settings
		needAuth.Get("/settings", h)
	}
	return app

}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	//第三个参数因为中间件所以不需要判断其值，否则不会执行此处理程序
	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

func main() {
	app := newApp()
	// open http://localhost:8080/admin
	app.Run(iris.Addr(":8080"))
}
