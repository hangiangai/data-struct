package middleware

import "github.com/kataras/iris"

func SetRequest(ctx iris.Context) {
	//打印日志
	ctx.Application().Logger().Infof("request info: path:%s method:%s", ctx.Path(), ctx.Method())
	//设置允许访问域名
	ctx.Header("Access-Control-Allow-Origin", "*")
	//设置允许的请求类型
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	//设置允许请求方法 DELETE
	ctx.Header("Access-Control-Allow-Methods", "DELETE")
	//设置允许请求方法 PUT
	ctx.Header("Access-Control-Allow-Methods", "PUT")
	//设置允许请求方法 POST
	ctx.Header("Access-Control-Allow-Methods", "POST")
	//设置数据类型
	ctx.ContentType("application/json")
	//向下执行
	ctx.Next()
}
