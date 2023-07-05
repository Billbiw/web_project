package model

//导入mysql驱动

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"v1/utils"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := utils.DbUser + ":" + utils.DbPassword + "@tcp(" + utils.DbHost + ":" + utils.DbPort + ")/" + utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用复数表名
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
	}
	//err := Db.AutoMigrate(&User{}, &Article{})
	//if err != nil {
	//	panic(err)
	//}

}
