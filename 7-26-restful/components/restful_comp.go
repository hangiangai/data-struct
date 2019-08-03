package components

import (
	"7-26-restful/model"
	util "7-26-restful/utils"
	"database/sql"

	"github.com/kataras/iris"
)

var (
	smt = []string{
		"select id,name,age,gender from student limit 10",
		"select id,name,age,gender from student where id=?",
		"select uid,u_email,u_gender,u_qq,u_tel from h_users",
	}
)

//参数接受对象
type Restful struct {
	db  *sql.DB //数据库对
	mid []iris.Handler
	mod map[string]model.Model
}

//组件
func (r Restful) Mount(cos Components) {
	r = Restful{
		db:  cos.Db,
		mod: cos.Mod,
		mid: cos.Mid,
	}
	cos.App.AllowMethods(iris.MethodOptions)
	router := cos.App.Party("/hangiangai/", r.mid...)
	//路由组
	router.Get("{tname:string}", r.GetHandle)
	router.Get("{tname:string}/{objectId:string}", r.GetByIdHandel)
	router.Delete("{tname:string}/{objectId:string}", r.DeleteHandel)
	router.Put("{tname:string}/{objectId:string}", r.PutHandle)
	router.Post("{tname:string}", r.PostHandle)
}

func (r Restful) GetHandle(ctx iris.Context) {

	tname := ctx.Params().Get("tname")
	rows, err := r.db.Query(smt[0])
	data := r.mod[tname].Data(rows)
	defer rows.Close()

	if err == nil {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  "get all data",
			"data": data,
		})
	} else {
		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "get all data",
			"data": data,
		})
	}

}

func (r Restful) DeleteHandel(ctx iris.Context) {

	tname := ctx.Params().Get("tname") //查询表名
	uid := ctx.Params().Get("uid")     //用户uid
	smt := util.Delete(tname, uid)     //生成sql语句
	result, err := r.db.Exec(smt)

	if err == nil { //smt执行成功
		affected, err := result.RowsAffected()
		if err != nil && affected > 0 {
			ctx.JSON(map[string]interface{}{
				"code": 404,
				"msg":  "delete failed",
				"data": map[string]interface{}{
					"uid": uid,
				},
			})
		} else {
			ctx.JSON(map[string]interface{}{
				"code": 404,
				"msg":  "delete failed",
				"data": map[string]interface{}{
					"uid": uid,
				},
			})
		}

	} else { //smt执行失败

		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "delete failed",
			"data": map[string]interface{}{
				"uid": uid,
			},
		})
	}
}

func (r Restful) PutHandle(ctx iris.Context) {

	uid := ctx.Params().Get("uid")
	tname := ctx.Params().Get("tname")
	del_data := make(map[string]string)
	ctx.ReadJSON(&del_data)
	smt := util.Update(tname, del_data, uid)
	result, err := r.db.Exec(smt)

	if err == nil {
		rows, err := result.RowsAffected()
		if err == nil && rows > 0 {
			ctx.JSON(map[string]interface{}{
				"code": 200,
				"msg":  "update successful",
				"data": map[string]string{
					"objectId": uid,
				},
			})
		} else {
			ctx.JSON(map[string]interface{}{
				"code": 404,
				"msg":  "update failed",
				"data": map[string]string{
					"objectId": uid,
				},
			})
		}
	} else {
		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "update failed",
			"data": map[string]string{
				"objectId": uid,
			},
		})
	}
}

func (r Restful) PostHandle(ctx iris.Context) {

	tname := ctx.Params().Get("tname")
	ins_data := make(map[string]string)
	ctx.ReadJSON(&ins_data)
	sql := util.Insert(tname, ins_data)
	result, err := r.db.Exec(sql)

	//数据库语句执行失败
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "insert the success",
			"data": "nil",
		})
		return
	}
	rows, err := result.RowsAffected()

	//数据库语句执行成功,但数据插入失败
	if err != nil || rows == 0 {
		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "insert failed",
			"data": "nil",
		})
		return
	}

	//插入成功
	ctx.JSON(map[string]interface{}{
		"code": 200,
		"msg":  "insert the success",
		"data": "nil",
	})

}

func (r Restful) GetByIdHandel(ctx iris.Context) {

	tname := ctx.Params().Get("tname")
	uid := ctx.Params().Get("uid")

	rows, err := r.db.Query(smt[1], uid)
	//数据库语句执行失败
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  "data acquisition success",
			"data": map[string]interface{}{
				"uid": uid,
			},
		})
		return
	}
	data := r.mod[tname].Data(rows)
	//未找到数据
	if len(data) == 0 {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  "this user does not exist",
			"data": map[string]string{
				"uid": uid,
			},
		})
		return
	}
	//返回查询的数据
	ctx.JSON(map[string]interface{}{
		"code": 200,
		"msg":  "this user does not exist",
		"data": data,
	})

}
