package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/duyanghao/gin-apiserver/pkg/models"
	"github.com/duyanghao/gin-apiserver/pkg/service"
	"github.com/gin-gonic/gin"
)

type SqlStudent_Controller interface {
	SqlGetID(c *gin.Context)
	SqlAdd(c *gin.Context)
	SqlUpdate(c *gin.Context)
	SqlDel(c *gin.Context)
}

// SqlStudent_Controller
// toDoController类型的struct实例 已经实现了该接口
func Sql_new_Controller() SqlStudent_Controller {
	return &sqlStudent_Controller{
		sqlClass1_Service: service.Sql_new_Service(),
	}
}

type sqlStudent_Controller struct {
	sqlClass1_Service service.SqlStudent_Service // struct
}

func (this *sqlStudent_Controller) SqlGetID(c *gin.Context) {
	fmt.Println("=================== controller: start SqlGetID")
	idParams := c.Param("id_Get")
	idParams_int, _ := strconv.Atoi(idParams)
	res := this.sqlClass1_Service.GetID_st(idParams_int)
	// c.JSON(200, models.Response{Code: 0, Message: res})
	c.JSON(200, res)

	return
}

func (this *sqlStudent_Controller) SqlUpdate(c *gin.Context) {
	fmt.Println("=================== controller: start SqlUpdate")
	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.JSON(-200, models.Response{Code: -1, Message: "error in Update"})
	}
	req_Upd := models.StudentClass1{}
	err = json.Unmarshal(reqBody, &req_Upd)
	if err != nil {
		fmt.Println(err)
	}

	this.sqlClass1_Service.Update_st(req_Upd)
	c.JSON(200, models.Response{Code: 0, Message: "update done"})

	return
}

func (this *sqlStudent_Controller) SqlAdd(c *gin.Context) {
	fmt.Println("=================== controller: start SqlAdd")
	reqBody, err := io.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.JSON(-200, models.Response{Code: -1, Message: "error in Add"})
	}
	req_Add := models.StudentClass1{}
	err = json.Unmarshal(reqBody, &req_Add)

	if err != nil {
		fmt.Println(err)
	}

	// 以下内容：调用POSTMAN时使用
	// req_Add := models.StudentClass1{}
	// tmp := c.Request.FormValue("age")
	// req_Add.StudentAge, _ = strconv.Atoi(tmp)
	// req_Add.StudentID, _ = strconv.Atoi(c.Request.FormValue("id"))
	// req_Add.StudentName = c.Request.FormValue("name")

	this.sqlClass1_Service.Add_st(req_Add)
	c.JSON(200, models.Response{Code: 0, Message: "add done"})

	return
}

func (this *sqlStudent_Controller) SqlDel(c *gin.Context) {
	fmt.Println("=================== controller: start SqlDel")
	idToDel := c.Param("id_Del")
	idToDel_int, _ := strconv.Atoi(idToDel)
	res := this.sqlClass1_Service.Del_st(idToDel_int)
	c.JSON(200, models.Response{Code: 0, Message: res, Data: 0})

	return
}
