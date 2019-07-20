package main

import (
	"fmt"
	"os"
	"time"
)

//create log
func createLog(user string, action string) string {
	//获取当前时间
	t := time.Now()
	//获取年月日
	year, month, day := t.Date()
	//获取时分秒
	hour, min, sec := t.Clock()
	dateString := fmt.Sprintf("%d/%d/%d %d:%d:%d\n", year, month, day, hour, min, sec)
	return user + " " + "'" + action + "'" + dateString
}

//写入日志
func Write(user string, action string) error {

	file, err := os.Open("logs.txt")
	if err != nil {
		fmt.Println("0")
		if os.IsNotExist(err) {
			file, err = os.Create("logs.txt")
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	_, err = file.WriteString("1234")

	defer file.Close()

	if err != nil {
		fmt.Println("4")
		fmt.Println(err)
		return err
	}

	return nil
}

func main() {

	file, err := os.Open("./root.log")
	if err != nil {
		fmt.Println(err, "22222222")
	}

	_, err = file.Write([]byte("123"))

	if err != nil {
		fmt.Println(err, "11111111")
	}

	_, err = file.Seek(0, 2)

	if err != nil {
		fmt.Println(err, "33333333")
	}

	_, err = file.Write([]byte("12345678"))

	if err != nil {
		fmt.Println(err, "444444444444444")
	}

	defer file.Close()

}

//用于绑定在api执行之前的操作(用于验证)
func Use() {

}
