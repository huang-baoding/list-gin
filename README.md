项目介绍：本项目为笔者在学习Gin框架的过程中写的一个清单小程序
使用到的技术栈有：Vue（前端）、Gin、Gorm、MySQL


部署运行：
运行环境：go 1.18及以上
需先在数据库中创建"list_gin"数据库；
将项目Clone到本地后：
(1)修改dao目录中的数据库连接配置(修改为自己的“用户：密码”)
(2)go mod tidy
(3)go run main.go
(4)在浏览器输入127.0.0.1:9090即可访问主页


目录结构：
static存放静态文件
templates存放模板文件
routers存放的是路由
controller存放控制器，控制代码路由进来之后执行什么函数
dao存放跟mysql数据库连接相关的内容
models存放模型以及对模型的增删改查
注：一般还有一个业务逻辑文件夹(logic)，存放操作逻辑，controller要改动数据时调用logic，logic再调用models的增删改查。但此项目较简单而不进行细分。


list-gin-old.zip为此项目在进行结构优化之前的版本
