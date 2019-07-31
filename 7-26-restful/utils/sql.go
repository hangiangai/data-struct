package util

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

//数据库配置文件
const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "hangiangai"
)

//连接数据库
func DatabaseConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	if db.Ping() != nil {
		log.Fatal(err.Error())
	}
	database = db
	return db
}

func Database() *sql.DB {
	return database
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
