package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //数据库驱动
)

var (
	DB *gorm.DB
)

//连接mysql数据库
func InitMySQL() (err error) {
	//连接用户为root，密码为fantastic918，连接的数据库为list_gin,默认字符类型为utf8mb4
	dsn := "root:fantastic0918@tcp(127.0.0.1:3306)/list_gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping() //测试连通性，联通了返回nil
}

func Close() {
	DB.Close()
}
