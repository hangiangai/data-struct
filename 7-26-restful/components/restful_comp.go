package components

import "github.com/kataras/iris"

//参数接受对象
type Restful struct {
}

//组件
func (r Restful) Mount(cos Components) {
	//创建路由组
	router := cos.App.Party("/api/v1/")

	{
		router.Get(cos.Custom["tablename"], func(ctx iris.Context) {
			ctx.WriteString(cos.Custom["tablename"])
		})

		router.Get(cos.Custom["tablename"], func(ctx iris.Context) {
			objectId := ctx.Params().Get("objectId")
			ctx.WriteString(objectId)
		})
	}

}
