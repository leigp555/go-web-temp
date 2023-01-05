package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"web/img/config"
)

var Mydb *gorm.DB //方便外面能使用到数据库
func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("root:%s@tcp(%s)/note?charset=utf8mb4&parseTime=True", config.Config.Mysql.Password, config.Config.Mysql.Addr)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置一个日志
	})
	if err != nil {
		return nil, err
	}
	Mydb = d
	fmt.Println("数据库连接成功")
	return Mydb, nil
}
