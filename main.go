package main

import (
	"fmt"
	"log"
	"web/img/database"
	"web/img/server"
)

func main() {
	//启动数据库连接
	_, err := database.InitDB()
	if err != nil {
		log.Panicf("数据库连接失败%v\n", err)
	} else {
		fmt.Println("成功连接数据库")
	}
	//启动服务器连接
	_, err = server.Serve()
	if err != nil {
		log.Panicf("服务开启失败%v\n", err)
	} else {
		fmt.Println("成功连接数据库")
	}

}
