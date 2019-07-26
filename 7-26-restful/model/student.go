package model

import (
	"database/sql"
	"fmt"
)

type Student struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Age    string `json:"age" db:"age"`
	Gender string `json:"gender" db:"gender"`
}

//用户重写
func (stu Student) GetData(rows *sql.Rows) []interface{} {
	result := make([]interface{}, 0, 0)
	for rows.Next() {
		//
		err := rows.Scan(&stu.Id, &stu.Name, &stu.Age, &stu.Gender)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, stu)
	}
	return result
}
