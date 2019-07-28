package restful

import (
	"database/sql"
	"fmt"
	"strings"

	"7-26-restful/model"

	"github.com/kataras/iris"
)

var (
	database    *sql.DB
	application *iris.Application
	accessUrl   = "/api/"
)

type Restful struct {
	Tablename string
	modelObj  model.Model
}

type Response struct {
	Message string
	Code    string
	Data    interface{}
}

//返回一个Restful对象
func New() *Restful {
	fmt.Println("1")
	return &Restful{}
}

//用于初始化一些数据
func (rf *Restful) Init(db *sql.DB, app *iris.Application) {
	fmt.Println("1")
	//初始化数据
	database = db
	application = app
	fmt.Println("12")
	UserRouter(app)

	fmt.Println("123")
}

//为指定表注册加restful风格的api接口
//model_
//tablename_
func (rf *Restful) Register(model_ model.Model, tablename_ string) {
	//创建一个新对象
	rf = &Restful{tablename_, model_}
	//处理OPTIONS预检
	application.AllowMethods(iris.MethodOptions)
	//创建路由组
	routes := application.Party(accessUrl, setRequestMiddleware)
	//restful风格路由
	{
		//获取全部数据
		routes.Get(tablename_, rf.GetData)
		//获取单条数据
		routes.Get(tablename_+"/{objectId:string}", rf.GetOneData)
		//添加数据
		routes.Post(tablename_, rf.InsertData)
		//修改数据
		routes.Put(tablename_+"/{objectId:string}", rf.UpdataData)
		//删除数据
		routes.Delete(tablename_+"/{objectId:string}", rf.DeleteData)
	}

}

//获取数据 需要用户实现
func (rf *Restful) GetData(ctx iris.Context) {

	rows, err := database.Query("select * from " + rf.Tablename)
	ErrProcess(err, 77)
	//用户实现这个方法
	data := rf.modelObj.GetData(rows)
	defer rows.Close()

	ctx.JSON(Response{
		Message: "Get all the data from the " + rf.Tablename,
		Code:    "200",
		Data:    data,
	})
}

//获取单条数据
func (rf *Restful) GetOneData(ctx iris.Context) {
	objectId := ctx.Params().Get("objectId")
	row, err := database.Query("select * from "+rf.Tablename+" where id=?", objectId)
	info, err := row.ColumnTypes()
	for _, val := range info {
		fmt.Println(val)
	}
	ErrProcess(err, 93)
	data := rf.modelObj.GetData(row)
	defer row.Close()
	if len(data) != 0 {
		ctx.JSON(Response{
			Message: "Get all " + rf.Tablename + " data",
			Code:    "200",
			Data:    data,
		})
	} else {
		ctx.JSON(Response{
			Message: "the user was not found",
			Code:    "404",
			Data:    nil,
		})
	}
}

//删除数据
func (rf *Restful) DeleteData(ctx iris.Context) {
	objectId := ctx.Params().Get("objectId")
	sql := CreateDeleteSql(rf.Tablename, objectId)
	res, err := database.Exec(sql)
	ErrProcess(err, 100)
	n, err := res.RowsAffected()
	ErrProcess(err, 102)
	if err == nil && n != 0 {
		ctx.JSON(Response{
			Message: "Data deleted successfully",
			Code:    "001",
			Data:    nil,
		})
	} else {
		ctx.JSON(Response{
			Message: "data deletion failed",
			Code:    "-002",
			Data:    nil,
		})
	}

}

//更新数据
func (rf *Restful) UpdataData(ctx iris.Context) {
	objectId := ctx.Params().Get("objectId")
	data := make(map[string]string)
	err := ctx.ReadJSON(&data)
	ErrProcess(err, 133)

	sql := CreateUpdateSql(rf.Tablename, data, objectId)
	fmt.Println(sql) //打印sql语句
	rows, err := database.Exec(sql)
	ErrProcess(err, 137)

	num, err := rows.RowsAffected()
	ErrProcess(err, 140)

	if err == nil && num != 0 { //返回错误或影响的行数不为1
		ctx.JSON(Response{
			Message: "data update successful",
			Code:    "001",
			Data:    nil,
		})
	} else {
		ctx.JSON(Response{
			Message: "data update failed",
			Code:    "-002",
			Data:    nil,
		})
	}
}

//插入数据
func (rf *Restful) InsertData(ctx iris.Context) {
	data := make(map[string]string)
	err := ctx.ReadJSON(&data)
	ErrProcess(err, 98)
	//打印添加数据
	sql := CreateInsertSql(rf.Tablename, data)
	//打印sql语句
	fmt.Println(sql)

	rows, err := database.Exec(sql)
	ErrProcess(err, 106)

	num, err := rows.RowsAffected()
	ErrProcess(err, 109)
	//返回数据
	if num != 0 {
		ctx.JSON(Response{
			Message: "data inserted successfully",
			Code:    "001",
			Data:    nil,
		})
	} else {
		ctx.JSON(Response{
			Message: "Insert the failure",
			Code:    "-002",
			Data:    nil,
		})
	}
}

//
func CreateInsertSql(tablename string, data map[string]string) string {
	keys := make([]string, 0, len(data))
	vals := make([]string, 0, len(data))
	for key, value := range data {
		keys = append(keys, key)
		vals = append(vals, "'"+value+"'")
	}
	sql := "insert into " + tablename + " (" + strings.Join(keys, ",") + ") values (" + strings.Join(vals, ",") + ")"
	return sql
}

//生成update sql语句
//tablename(表名) data(更新数据) objectId(更新对象)
func CreateUpdateSql(tablename string, data map[string]string, objectId string) string {
	param := make([]string, 0, len(data))
	for key, value := range data {
		param = append(param, key+"="+"'"+value+"'")
	}
	sql := "update " + tablename + " set " + strings.Join(param, ",") + " where id = " + objectId

	return sql
}

//delete sql
// tablename(表名) objectId(删除对象)
func CreateDeleteSql(tablename string, objectId string) string {
	return "delete from " + tablename + " where id=" + objectId
}

//错误处理
func ErrProcess(err error, n int) {
	if err != nil {
		fmt.Printf("error = %s %d\n", err.Error(), n)
	}
}
