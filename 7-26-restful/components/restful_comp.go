package components

import (
	"7-26-restful/model"
	util "7-26-restful/utils"
	"database/sql"
	"fmt"

	"github.com/kataras/iris"
)

var (
	smt = []string{
		"select id,name,age,gender from student limit 10",
		"select id,name,age,gender from student where id=?",
	}
)

//参数接受对象
type Restful struct {
	database  *sql.DB
	model     []model.Model
	tablename string
}

//组件
func (r Restful) Mount(cos Components) {

	r = Restful{
		database:  cos.Db,
		model:     cos.Mod,
		tablename: cos.Tname,
	}

	cos.App.AllowMethods(iris.MethodOptions)

	//创建路由组
	router := cos.App.Party("/api/v1/", cos.Mid...)

	{
		router.Get(r.tablename, r.getAllData)

		router.Get(r.tablename+"/{objectId:string}", r.getByObjectId)

		router.Delete(r.tablename+"/{objectId:string}", r.deleteData)

		router.Put(r.tablename+"/{objectId:string}", r.updateData)

		router.Post(r.tablename, r.insertData)
	}

}

func (r Restful) getAllData(ctx iris.Context) {
	rows, err := r.database.Query(smt[0])
	if err != nil {
		fmt.Println(err)
	}
	data := r.model[0].GetData(rows)
	defer rows.Close()

	ctx.JSON(map[string]interface{}{
		"code": 200,
		"msg":  "get all data",
		"data": data,
	})

}

func (r Restful) getByObjectId(ctx iris.Context) {
	objectId := ctx.Params().Get("objectId")
	rows, err := r.database.Query(smt[1], objectId)
	if err != nil {
		fmt.Println(err)
	}
	data := r.model[0].GetData(rows)
	if len(data) > 0 {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  objectId,
			"data": data,
		})
	} else {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  "this user does not exist",
			"data": map[string]string{
				"objectId": objectId,
			},
		})
	}

}

func (r Restful) deleteData(ctx iris.Context) {
	objectId := ctx.Params().Get("objectId")
	sql := "delete from " + r.tablename + " where id=?"
	result, err := r.database.Exec(sql, objectId)
	if err != nil {
		fmt.Println(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowsAffected >= 0 {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  "delete the success",
			"data": map[string]string{
				"objectId": objectId,
			},
		})
	} else {
		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "delete failed",
			"data": map[string]string{
				"objectId": objectId,
			},
		})
	}
}

func (r Restful) updateData(ctx iris.Context) {
	objectId := ctx.Params().Get("objectId")
	delete_data := make(map[string]string)
	ctx.ReadJSON(&delete_data)
	sql := util.CreateUpdateSql(r.tablename, delete_data, objectId)
	result, err := r.database.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowsAffected >= 0 {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  "update the success",
			"data": map[string]string{
				"objectId": objectId,
			},
		})
	} else {
		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "update failed",
			"data": map[string]string{
				"objectId": objectId,
			},
		})
	}
}

func (r Restful) insertData(ctx iris.Context) {

	insert_data := make(map[string]string)
	ctx.ReadJSON(&insert_data)
	sql := util.CreateInsertSql(r.tablename, insert_data)
	result, err := r.database.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowsAffected >= 0 {
		ctx.JSON(map[string]interface{}{
			"code": 200,
			"msg":  "insert the success",
			"data": "nil",
		})
	} else {
		ctx.JSON(map[string]interface{}{
			"code": 404,
			"msg":  "insert failed",
			"data": "nil",
		})
	}
}
