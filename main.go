package main

import (
	"log"

	"github.com/huang-baoding/list-gin/dao"
	"github.com/huang-baoding/list-gin/models"
	"github.com/huang-baoding/list-gin/routers"
)

func main() {

	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		log.Panic(err)
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	models.InitModel()

	//路由
	r := routers.SetupRouter()
	r.Run(":9090")
}
