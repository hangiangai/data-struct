package model

import (
	"database/sql"
)

//数据模型
type Model interface {
	//用户必需实现的接口
	GetData(rows *sql.Rows) []interface{} //获取数据
}

//响应模型
type Response struct {
	Code    string
	Message string
	Data    interface{}
	Token   string
}
