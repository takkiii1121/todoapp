package main

import (
	"awesomeProject/TODO APP/Config"
	"awesomeProject/TODO APP/Models"
	"awesomeProject/TODO APP/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	//DBとの連携
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	//エラーがあれば表示
	if err != nil {
		fmt.Println("status:", err)
	}

	defer Config.DB.Close()
	//自動でマイグレーションする。
	//ただし、一度作ったカラム、インデックスは修正できない
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	//デフォルトは8080
	r.Run()
}
