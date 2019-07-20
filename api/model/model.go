package dbModel

import (
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	ParamsAddr []interface{} //存储用于获取数据库数据的关键字 以地址形式存储
	ParamsStr  []string      //存储用于获取数据库数据的关键字 以字符串形式存储
	DbModel    interface{}   //数据模型
}

//初始化模型
func Init(model interface{}) *Model {
	//通过反射获取model的信息
	getVal := reflect.ValueOf(model).Elem()
	getType := reflect.TypeOf(model).Elem()
	length := getType.NumField()
	paramsAddr := make([]interface{}, length)
	paramsStr := make([]string, length)
	for i := 0; i < length; i++ {
		paramsAddr[i] = getVal.Field(i).Addr().Interface()
		paramsStr[i] = strings.ToLower(getType.Field(i).Name)
	}
	return &Model{paramsAddr, paramsStr, model}
}

//生成查询语句
//column(查询关键字) tablename(表名)
func CreateQuerySql(column []string, tablename string) string {
	return "select " + strings.Join(column, ",") + " from " + tablename + " limit 10"
}

func CreateQueryByIdSql(column []string, tablename string, objectId string) string {
	return "select " + strings.Join(column, ",") + " from " + tablename + " where id=" + objectId
}

//delete sql
// tablename(表名) objectId(删除对象)
func CreateDeleteSql(tablename string, objectId string) string {
	return "delete from " + tablename + " where id=" + objectId
}

//生成insert sql语句
//tablename(表名) data(插入数据)
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

//响应模型
type ResponseModel struct {
	Message string      //响应消息
	Code    string      //响应状态码
	Result  interface{} //返回结果
}
