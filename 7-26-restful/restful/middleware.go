package restful

import "github.com/kataras/iris"

//设置请求信息
func setRequestMiddleware(ctx iris.Context) {
	//打印日志
	ctx.Application().Logger().Infof("request info: path:%s method:%s", ctx.Path(), ctx.Method())
	//设置允许访问域名
	ctx.Header("Access-Control-Allow-Origin", "*")
	//设置允许请求方法 DELETE
	ctx.Header("Access-Control-Allow-Methods", "DELETE")
	//设置允许请求方法 PUT
	ctx.Header("Access-Control-Allow-Methods", "PUT")
	//设置数据类型
	ctx.ContentType("application/json")
	//向下执行
	ctx.Next()
}
