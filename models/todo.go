package models

import "github.com/huang-baoding/list-gin/dao"

// 数据库连接模型
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func InitModel() {
	//模型迁移，会在mysql中创建名为todos的表
	dao.DB.AutoMigrate(Todo{})
}

/*
	对Todo模型进行的增删改查
*/

//在mysql的list_gin数据库的todos表中创建一条记录
func CreateAtodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

//查询全部记录
func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

//查询一条记录
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

//更新一条记录
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

//删除一条记录
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
