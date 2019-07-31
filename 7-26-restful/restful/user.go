package restful

import (
	"fmt"

	"github.com/kataras/iris"
)

func UserRouter(app *iris.Application) {

	app.Get("/a", func(ctx iris.Context) {
		ctx.WriteString("1234")
	})

	fmt.Println("234")

}
