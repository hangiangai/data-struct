package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "hangiangai"
)

//base64加密
func Base64Encode(v interface{}) (string, error) {
	json, err := json.Marshal(v) //转化成json
	base64 := base64.StdEncoding.EncodeToString(json)
	return base64, err
}

//base64解密
func Base64Decode(v string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(v)
}

func main() {
	// app := iris.New()

	header := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}

	fmt.Println(Base64Encode(header))

	// payload := map[string]interface{}{
	// 	"iss": "hangiangai",              //签发者
	// 	"sub": "all",                     //登录用户名
	// 	"aud": "client",                  //登录端
	// 	"exp": time.Now().Unix() + 30*60, //过期时间
	// }

	fmt.Println(time.Now().Unix())

	fmt.Println()

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}

	row := db.QueryRow("select name from student where id = ?", 16)
	var name string
	row.Scan(&name)

	fmt.Println(name)
}
