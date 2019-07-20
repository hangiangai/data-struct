package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	//数据库配置参数
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "hangiangai"

	//api
	BaseUrl = "/api/"
)

//连接数据库
func ConnectDatabase() (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn) //
	if err != nil {
		return nil, err
	}
	err = db.Ping() //检测是否连接正常
	if err != nil {
		return nil, err
	}
	return db, nil
}
