package controller

import (
	"github.com/duyanghao/gin-apiserver/pkg/models"
	"github.com/duyanghao/gin-apiserver/pkg/service"
	"github.com/gin-gonic/gin"
)

type ToDoController interface {
	GetToDo(c *gin.Context) // 查找
}

// 普通函数，返回类型是ToDoController接口，
// toDoController类型的struct实例 已经实现了该接口
func NewToDoController() ToDoController {
	return &toDoController{
		toDoService: service.NewToDoService(), // toDoController结构体的成员
	}
}

type toDoController struct {
	toDoService service.ToDoService // service.ToDoService 是service中的一个struct
}

// @Summary GetToDo
// @Description GetToDo
// @Success 200 {object} models.Response OK
// @Failure 400 {object} models.Response Bad Request
// @Failure 401 {object} models.Response Unauthorized
// @Failure 403 {object} models.Response Forbidden
// @Failure 500 {object} models.Response Internal Server Error
// @router /todo/get [get]

// toDoController结构体 实现了 接口ToDoController的GetToDo方法
func (this *toDoController) GetToDo(c *gin.Context) {
	nameParams := c.Param("name")
	res := this.toDoService.Get(nameParams)
	// 按照name查询，返回查询结果res
	c.JSON(200, models.Response{Code: 0, Message: "todo demo, Data is the student's ID", Data: res})
	return
}
