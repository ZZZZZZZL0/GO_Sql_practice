package route

import (
	"os"

	_ "github.com/duyanghao/gin-apiserver/docs"
	"github.com/duyanghao/gin-apiserver/pkg/controller"

	//"github.com/duyanghao/gin-apiserver/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// @title Swagger gin-apiserver
// @version 0.1.0
// @description This is a gin-apiserver.
// @contact.name duyanghao
// @contact.url https://duyanghao.github.io
// @contact.email 1294057873@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func InstallRoutes(r *gin.Engine) {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	c := controller.Sql_new_Controller()
	c_name := controller.NewToDoController()

	// ---------------------------------------------------------
	r.GET("/names/:name", c_name.GetToDo)
	r.POST("/add", c.SqlAdd)       // http://192.168.1.1:8806/api/v1/users/add
	r.GET("/get/:id", c.SqlGetID)  // http://192.168.1.1:8806/api/v1/users/get/5
	r.POST("/update", c.SqlUpdate) // http://192.168.1.1:8806/api/v1/users/update
	r.POST("/del", c.SqlDel)       // http://192.168.1.1:8806/api/v1/users/del
	// ---------------------------------------------------------

	// a ping api test
	r.GET("/ping", controller.Ping)

	// get gin-apiserver version
	r.GET("/version", controller.Version)

	// config reload
	r.Any("/-/reload", func(c *gin.Context) {
		//log.Info("===== Server Stop! Cause: Config Reload. =====")
		os.Exit(1)
	})

	rootGroup := r.Group("/api/v1")
	{
		// a ping api to test basic auth
		rootGroup.GET("/ping", controller.Ping)
	}

	{
		sql_toDoController := controller.Sql_new_Controller()
		rootGroup.POST("/add", sql_toDoController.SqlAdd)
	}

	{
		sql_toDoController := controller.Sql_new_Controller()
		rootGroup.POST("/del/:id_Del", sql_toDoController.SqlDel)
	}

	{
		sql_toDoController := controller.Sql_new_Controller()
		rootGroup.POST("/update", sql_toDoController.SqlUpdate)
	}

	{
		sql_toDoController := controller.Sql_new_Controller()
		rootGroup.GET("/get/:id_Get", sql_toDoController.SqlGetID)
	}

	{
		toDoController := controller.NewToDoController()
		rootGroup.GET("/names/:name", toDoController.GetToDo)
	}
}
