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

//SqlSmt
type SqlUtil struct {
	smt bytes.Buffer
	Db  *sql.DB
}

func New() *SqlUtil {
	return &SqlUtil{
		Db: DatabaseConnection(),
	}
}

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
func (s *SqlUtil) Select(sel []string) *SqlUtil {
	var sct []string
	s.smt.WriteString(" select ")
	for _, val := range sel {
		sct = append(sct, "`"+val+"`")
	}
	s.smt.WriteString(strings.Join(sct, ","))
	return s
}

//生成where语句
func (s *SqlUtil) Where(key string, value string) *SqlUtil {
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
func (s *SqlUtil) From(key string) *SqlUtil {
	s.smt.WriteString(" from ")
	s.smt.WriteByte('`')
	s.smt.WriteString(key)
	s.smt.WriteByte('`')
	return s
}

//生成Limit语句
func (s *SqlUtil) Limit() {

}

//生成最终sql语句
func (s *SqlUtil) Smt() string {
	return s.smt.String()
}
