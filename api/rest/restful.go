package rest

import (
	"database/sql"
	"fmt"
	"goworkerspace/api/config"
	dbModel "goworkerspace/api/model"
	"log"

	"github.com/kataras/iris"
)

//config
var (
	database    *sql.DB                   //数据库对象
	accessUrl   string                    //基础url
	application *iris.Application         //
	isinit      bool              = false //是否初始化
)

type RESTful struct {
	TableName string         //表名
	Model     *dbModel.Model //模型
}

//初始化并返回一个RESTful对象
//app(必须) db,baseUrl(可选,不存在时会调用config文件的定义)
func New(app *iris.Application, args ...interface{}) *RESTful {

	isinit = true //已初始化
	application = app

	if len(args) > 0 {
		if db, ok := args[0].(*sql.DB); ok { //可选
			database = db
		}
		accessUrl = config.BaseUrl //只有一个参数时
	} else if len(args) > 1 {
		if url, ok := args[1].(string); ok {
			accessUrl = url
		}
	} else {
		db, err := config.ConnectDatabase()
		if err != nil {
			log.Fatal(err)
		}
		database = db
		accessUrl = config.BaseUrl
	}

	return &RESTful{}
}

//需要先调用InitParam 返回一个RESTful对象
//tablename(表名) model(模型)
func Register(tableName string, model *dbModel.Model) *RESTful {
	return &RESTful{
		TableName: tableName,
		Model:     model,
	}
}

func mix(fun1, fun2, string) string{
	return fun1(fun2(sting))
}

//添加RESTful Api
func (rf *RESTful) Add(restful *RESTful) {

	if !isinit { //是否初始化参数
		log.Fatal("uninitialized parameters,call the new method")
		return
	}

	rf = restful //初始化RESTful

	application.PartyFunc(accessUrl, func(query iris.Party) {
		query.Get(rf.TableName, rf.getAllData)
		query.Get(rf.TableName+"/{objectId:string}", rf.getDataById)
		query.Delete(rf.TableName+"/{objectId:string}", rf.deleteDataById)
		query.Post(rf.TableName, rf.addData)
		query.Put(rf.TableName+"/{objectId:string}", rf.updateDataById)
	})

}

func (rf *RESTful) getAllData(ctx iris.Context) { //Get 获取全部数据

	ctx.ContentType("application/json")

	sql := dbModel.CreateQuerySql(rf.Model.ParamsStr, rf.TableName)
	rows, err := database.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	result := make([]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(rf.Model.ParamsAddr...)
		if err != nil {
			fmt.Println("row.Scan err = ", err)
		}
		result = append(result, rf.Model.DbModel)
	}

	err = rows.Err() //检查是否全部遍历
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(dbModel.ResponseModel{Message: "successful", Code: "001", Result: result})
}

func (rf *RESTful) getDataById(ctx iris.Context) {

	ctx.ContentType("application/json")
	objectId := ctx.Params().Get("objectId")
	sql := dbModel.CreateQueryByIdSql(rf.Model.ParamsStr, rf.TableName, objectId)
	row := database.QueryRow(sql)
	err := row.Scan(rf.Model.ParamsAddr...)
	if err != nil {
		fmt.Println("row.Scan err = ", err)
		ctx.JSON(dbModel.ResponseModel{Message: "failed", Code: "002", Result: nil})
	} else {
		ctx.JSON(dbModel.ResponseModel{Message: "successful", Code: "001", Result: rf.Model.DbModel})
	}
}

func (rf *RESTful) deleteDataById(ctx iris.Context) {

	objectId := ctx.Params().Get("objectId")

	sql := dbModel.CreateDeleteSql(rf.TableName, objectId)
	rows, err := database.Exec(sql)
	if err != nil {
		log.Fatal("delete data err", err.Error())
	}

	num, err := rows.RowsAffected() //影响的行数
	if err != nil && num == 1 {
		ctx.JSON("Success") //返回被删除的数据
	} else {
		ctx.JSON(`{"code:":""}`)
	}
}

//添加数据
func (rf *RESTful) addData(ctx iris.Context) {

	ctx.ContentType("application/json")

	data := make(map[string]string)

	err := ctx.ReadJSON(&data)

	fmt.Println(data)
	if err != nil {
		fmt.Println("ctx.ReadJSON err = ", err)
	}

	sql := dbModel.CreateInsertSql(rf.TableName, data)

	fmt.Println(sql)
	rows, err := database.Exec(sql)
	if err != nil {
		fmt.Println("api.Database.Exec err = ", err)
	}

	num, err_ := rows.RowsAffected()
	if err_ != nil {
		fmt.Println("rows.RowsAffected err = ", err_)
	}
	fmt.Println(num)

	ctx.JSON(dbModel.ResponseModel{Message: "successful", Code: "001", Result: data})

}

//更新数据
func (rf *RESTful) updateDataById(ctx iris.Context) {

	objectId := ctx.Params().Get("objectId")

	data := make(map[string]string)
	err := ctx.ReadJSON(&data)
	if err != nil {
		fmt.Println("ctx.ReadJSON err = ", err)
	}

	sql := dbModel.CreateUpdateSql(rf.TableName, data, objectId)
	rows, err := database.Exec(sql)
	if err != nil {
		log.Println("api.Database.Exec err = ", err.Error())
	}

	num, err_ := rows.RowsAffected()
	if err_ != nil || num != 1 { //返回错误或影响的行数不为1
		fmt.Println("rows.RowsAffected err = ", err_)
		ctx.JSON(dbModel.ResponseModel{Message: "successful", Code: "002", Result: data})
	} else {
		ctx.JSON(dbModel.ResponseModel{Message: "successful", Code: "001", Result: data})
	}
}
