package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/huang-baoding/list-gin/controller"
)

func SetupRouter() *gin.Engine {
	//创建默认路由
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找,前一个参数是当前项目的哪个目录，
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	//把模板返回并渲染给浏览器
	r.GET("/", controller.IndexHandler)

	// v1路由组
	v1Group := r.Group("v1")
	{
		// 添加待办事项
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
