package service

import (
	"ImApplication.go/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngine *xorm.Engine

func init() {
	var err error // 声明错误变量
	// 初始化xorm引擎
	DbEngine, err = xorm.NewEngine("mysql", "root:qm36TMOFFPHmyoaE@tcp(192.168.123.101:3306)/im?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 设置日志输出，显示SQL语句
	DbEngine.ShowSQL(true)

	// 自动创建表结构
	err = DbEngine.Sync2(new(model.User))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("database connect succect")
}
