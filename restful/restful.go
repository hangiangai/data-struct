package restful

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/kataras/iris"
)

var (
	application *iris.Application
	database    *sql.DB
	accessUrl   = "/api/"
)

//
type Restful struct {
	Tablename string //数据库表明
}

type Response struct {
	Message string
	Code    string
	Data    interface{}
}

//创建对象
func New(app *iris.Application, db *sql.DB) *Restful {
	//添加数据库对象
	database = db
	application = app
	return &Restful{}
}

//注册
func (restful *Restful) Register(tablename string) {

	restful = &Restful{Tablename: tablename}
	//允许options预检
	application.AllowMethods(iris.MethodOptions)

	routes := application.Party(accessUrl, setRequestMiddleware)
	//restful风格路由
	{
		//获取全部数据
		routes.Get(tablename, restful.getAllData)
		//获取单条数据
		routes.Get(tablename+"/{objectId:string}", restful.getOneData)
		//添加数据
		routes.Post(tablename, restful.insertData)
		//修改数据
		routes.Put(tablename+"/{objectId:string}", restful.updateData)
		//删除数据
		routes.Delete(tablename+"/{objectId:string}", restful.deleteData)
	}

}

//获取10条数据
func (restful *Restful) getAllData(ctx iris.Context) {
	//访问记录
	rows, err := database.Query("select * from " + restful.Tablename)
	ErrProcess(err, 58) //错误处理函数
	defer rows.Close()
	result := rowsProcess(rows)
	//返回数据
	ctx.JSON(Response{Message: "successful", Code: "001", Data: result})

}

//获取指定id用户
func (restful *Restful) getOneData(ctx iris.Context) {

	objectId := ctx.Params().Get("objectId")
	//sql语句
	sql := "select * from " + restful.Tablename + " where id= '" + objectId + "'"
	fmt.Println(sql)
	rows, err := database.Query(sql)
	ErrProcess(err, 75)
	result := rowsProcess(rows)

	if len(result) > 0 {
		ctx.JSON(Response{
			Message: "successful",
			Code:    "001",
			Data:    result,
		})
	} else {
		ctx.JSON(Response{
			Message: "user does not exist",
			Code:    "-002",
			Data:    result,
		})
	}
}

//插入数据
func (restful *Restful) insertData(ctx iris.Context) {

	data := make(map[string]string)
	err := ctx.ReadJSON(&data)
	ErrProcess(err, 98)
	//打印添加数据
	sql := CreateInsertSql(restful.Tablename, data)
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

//更新数据
func (restful *Restful) updateData(ctx iris.Context) {

	objectId := ctx.Params().Get("objectId")
	data := make(map[string]string)
	err := ctx.ReadJSON(&data)
	ErrProcess(err, 133)

	sql := CreateUpdateSql(restful.Tablename, data, objectId)
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

	ctx.WriteString("123")
}

//删除数据
func (restful *Restful) deleteData(ctx iris.Context) {

	objectId := ctx.Params().Get("objectId")
	sql := CreateDeleteSql(restful.Tablename, objectId)
	fmt.Println(sql) //打印sql语句
	rows, err := database.Exec(sql)
	ErrProcess(err, 168)
	num, err := rows.RowsAffected() //影响的行数
	if err == nil && num != 0 {
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

//处理rows数据
func rowsProcess(rows *sql.Rows) []map[string]string {

	columns, err := rows.Columns()
	ErrProcess(err, 118) //错误处理

	length := len(columns)
	//存储每次循环的值
	values := make([]interface{}, length)
	//存储values每个对应值得地址
	scanParam := make([]interface{}, length)
	for i := range values {
		scanParam[i] = &values[i]
	}
	//存储所有数据
	res := make([]map[string]string, 0)
	//遍历rows
	for rows.Next() {
		//用于存储一条数据
		onedata := make(map[string]string)
		err = rows.Scan(scanParam...)
		ErrProcess(err, 135)
		for i := range values {
			//存储数据
			onedata[columns[i]] = string(values[i].([]byte))
		}
		res = append(res, onedata)
	}
	return res
}

//错误处理
func ErrProcess(err error, n int) {
	if err != nil {
		fmt.Printf("error = %s %d\n", err.Error(), n)
	}
}

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
