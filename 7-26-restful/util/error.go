package util

import "fmt"

//对错误进行简单处理 返回错误的信息和错误出现所在行
func ErrProcess(err error, n int) {
	if err != nil {
		fmt.Printf("error = %s %d\n", err.Error(), n)
	}
}
