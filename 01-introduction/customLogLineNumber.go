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
	defer f.Close()

	// 日期和源文件名
	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstdFlags)
	iLog.Println("Go language")
	// 重新设置标志格式
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry!")
}
