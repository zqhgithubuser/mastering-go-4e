package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "go.log")
	fmt.Println(LOGFILE)

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// 函数返回之前关闭文件
	defer f.Close()
	// 日志条目前缀
	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there!")
}
