package util

import (
	"bytes"
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

//SqlSmt
type Sql struct {
	smt bytes.Buffer
}

//创建插入语句
func Insert(tname string, data map[string]string) string {
	var smt bytes.Buffer
	var field_ []string
	fmt.Println(field_)
	var value_ []string
	fmt.Println(value_)
	smt.WriteString("insert into ")
	smt.WriteByte('`')
	smt.WriteString(tname)
	smt.WriteByte('`')
	smt.WriteByte(' ')
	//处理插入数据
	for k, v := range data {
		var kfiled bytes.Buffer
		kfiled.WriteByte('`')
		kfiled.WriteString(k)
		kfiled.WriteByte('`')
		field_ = append(field_, kfiled.String())
		var vfiled bytes.Buffer
		vfiled.WriteByte('"')
		vfiled.WriteString(v)
		vfiled.WriteByte('"')
		value_ = append(value_, vfiled.String())
	}
	smt.WriteByte('(')
	smt.WriteString(strings.Join(field_, ","))
	smt.WriteByte(')')
	smt.WriteByte(' ')
	smt.WriteString("values")
	smt.WriteByte(' ')
	smt.WriteByte('(')
	smt.WriteString(strings.Join(value_, ","))
	smt.WriteByte(')')
	return smt.String()
}

func Delete(tname string, objectId string) string {
	var smt bytes.Buffer
	smt.WriteString("delete from ")
	smt.WriteByte('`')
	smt.WriteString(tname)
	smt.WriteByte('`')
	smt.WriteString(" where id=")
	smt.WriteByte('"')
	smt.WriteString(objectId)
	smt.WriteByte('"')
	return smt.String()
}

func Update(tname string, data map[string]string, oid string) string {
	var smt bytes.Buffer
	smt.WriteString("update ")
	smt.WriteString(tname)
	smt.WriteString(" set ")
	for k, v := range data {
		smt.WriteByte('`')
		smt.WriteString(k)
		smt.WriteByte('`')
		smt.WriteByte('=')
		smt.WriteByte('"')
		smt.WriteString(v)
		smt.WriteByte('"')
	}
	smt.WriteString(" where id=")
	smt.WriteByte('"')
	smt.WriteString(oid)
	smt.WriteByte('"')
	return smt.String()
}

//生成select语句
func (s *Sql) Select(sel []string) *Sql {
	var sct []string
	s.smt.WriteString(" select ")
	for _, val := range sel {
		sct = append(sct, "`"+val+"`")
	}
	s.smt.WriteString(strings.Join(sct, ","))
	return s
}

//生成where语句
func (s *Sql) Where(key string, value string) *Sql {
	s.smt.WriteString(" where ")
	s.smt.WriteByte('`')
	s.smt.WriteString(key)
	s.smt.WriteByte('`')
	s.smt.WriteByte('=')
	s.smt.WriteByte('"')
	s.smt.WriteString(value)
	s.smt.WriteByte('"')
	return s
}

//生成From语句
func (s *Sql) From(key string) *Sql {
	s.smt.WriteString(" from ")
	s.smt.WriteByte('`')
	s.smt.WriteString(key)
	s.smt.WriteByte('`')
	return s
}

//生成最终sql语句
func (s *Sql) Smt() string {
	return s.smt.String()
}
