package components

import (
	"7-26-restful/middleware"
	"7-26-restful/model"
	"database/sql"
	"log"
	"strings"

	"github.com/kataras/iris"
)

var (
	//存放组件
	component = map[string]Component{
		"user":    User{},
		"restful": Restful{},
	}
	datamodel = map[string]model.Model{
		"student": model.Student{},
		"teacher": model.Teacher{},
	}
	middlewares = map[string]iris.Handler{
		"token":   middleware.CheckToken,
		"request": middleware.SetRequest,
	}
)

//组件接口
type Component interface {
	Mount(cts Components)
}

//组件参数
type Components struct {
	//数据模型
	Mod []model.Model
	//iris对象
	App *iris.Application
	//database对象
	Db *sql.DB
	//存放组件所需中间件
	Mid []iris.Handler
	//数据表名称
	Tname string
	//Cus
	Cus []interface{}
}

/*
 *初始化并且返回Components对象
 *@method New()
 */

func New(db *sql.DB, app *iris.Application) *Components {
	return &Components{
		App: app,
		Db:  db,
		Mid: make([]iris.Handler, 0, 0),
		Mod: make([]model.Model, 0, 0),
		Cus: make([]interface{}, 0, 0),
	}
}

/*
 *使用注册的组件
 *@method Use()
 *@param:pen_name string 组件名称
 *@param:custom string 组件所需参数
 */

func (cos Components) Use(pen_name string, args map[string]string) {
	//组件存在
	if _, ok := component[pen_name]; !ok {
		log.Fatalf("the %s component does not exist", pen_name)
	}
	//处理参数
	for key, val := range args {
		switch key {
		case "mid": //中间件以|分割
			mid_ := strings.Split(val, "|")
			for _, val := range mid_ {
				cos.Mid = append(cos.Mid, middlewares[val])
			}
		case "mod": //处理模型
			mod_ := strings.Split(val, "|")
			for _, val := range mod_ {
				cos.Mod = append(cos.Mod, datamodel[val])
			}
		case "tname": //数据表名
			cos.Tname = val
		default: //用户传的指定数据
			cos.Cus = append(cos.Cus, val)
		}
	}
	//组件调用
	component[pen_name].Mount(cos)
}
