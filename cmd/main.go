package main

import (
	"fmt"

	"github.com/duyanghao/gin-apiserver/pkg/route"
	// "github.com/duyanghao/gin-apiserver/pkg/util"
	_ "github.com/duyanghao/gin-apiserver/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main start")

	// util.SetupSigusr1Trap()
	r := gin.Default()
	//m := config.GetString(config.FLAG_KEY_GIN_MODE)
	gin.SetMode("release")

	route.InstallRoutes(r)
	//serverBindAddr := fmt.Sprintf("%s:%d", config.GetString(config.FLAG_KEY_SERVER_HOST), config.GetInt(config.FLAG_KEY_SERVER_PORT))
	serverBindAddr := fmt.Sprintf("%s:%d", "0.0.0.0", 8080)
	//log.Infof("Run server at %s", serverBindAddr)

	r.Run(serverBindAddr) // listen and serve

}
