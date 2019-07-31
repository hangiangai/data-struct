package components

import (
	"7-26-restful/model"
	"database/sql"

	"github.com/kataras/iris"
)

var (
	//存放组件
	component = map[string]Component{
		"user":    User{},
		"restful": Restful{},
	}
)

//组件接口
type Component interface {
	Mount(cts Components)
}

//组件参数
type Components struct {
	//存放所有的数据模型
	Model map[string]model.Model
	//iris对象
	App *iris.Application
	//database对象
	Db *sql.DB
	//存放所有组件
	mount map[string]Component
	//除了组件名参数的其他参数
	Custom map[string]string
}

/*
 *初始化并且返回Components对象
 *@method New()
 */

func New(db *sql.DB, app *iris.Application) *Components {
	return &Components{
		App:   app,
		Db:    db,
		mount: component,
	}
}

/*
 *使用注册的组件
 *@method Use()
 *@param:cpname string 组件名称
 *@param:custom string 自定义数据
 */

func (cos *Components) Use(cpname string, custom map[string]string) {
	cos.Custom = custom
	if comp, ok := cos.mount[cpname]; ok {
		comp.Mount(*cos)
	}
}
