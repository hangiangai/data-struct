package model

import (
	"database/sql"
	"fmt"
)

type Teacher struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  string `json:"age" db:"age"`
}

//用户重写
func (tea Teacher) GetData(rows *sql.Rows) []interface{} {

	result := make([]interface{}, 0, 0)
	for rows.Next() {
		err := rows.Scan(&tea.Id, &tea.Name, &tea.Age)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, tea)
	}
	return result
}
