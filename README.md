gin-apiserver
===================

gin-apiserver is a HTTP apiserver framework based on [Gin](https://github.com/gin-gonic/gin).

## Introduction

`gin-apiserver`是一个基于[gin](https://github.com/gin-gonic/gin)框架写的ApiServer框架，主要用于企业生产环境中的快速开发

## Features

* 1、支持configmap reload api
  ```go
  // config reload
  r.Any("/-/reload", func(c *gin.Context) {
          log.Info("===== Server Stop! Cause: Config Reload. =====")
          os.Exit(1)
  })
  ```

* 2、支持ping-pong健康检查&版本获取
  ```go
  // a ping api test
  r.GET("/ping", controller.Ping)
  
  // get gin-apiserver version
  r.GET("/version", controller.Version)
  ```

* 3、支持dump-goroutine-stack-traces
  ```bash
  $ kill -SIGUSR1 41307
  
  === BEGIN goroutine stack dump ===
  goroutine 20 [running]:
  github.com/duyanghao/gin-apiserver/pkg/util.dumpStacks()
          /root/go/src/github.com/duyanghao/gin-apiserver/pkg/util/trap.go:23 +0x6d
  github.com/duyanghao/gin-apiserver/pkg/util.SetupSigusr1Trap.func1(0xc000332240)
          /root/go/src/github.com/duyanghao/gin-apiserver/pkg/util/trap.go:16 +0x34
  created by github.com/duyanghao/gin-apiserver/pkg/util.SetupSigusr1Trap
          /root/go/src/github.com/duyanghao/gin-apiserver/pkg/util/trap.go:14 +0xab
  
  goroutine 1 [IO wait]:
  internal/poll.runtime_pollWait(0x7fccf3b86f68, 0x72, 0x0)
          /usr/local/go/src/runtime/netpoll.go:182 +0x56
  internal/poll.(*pollDesc).wait(0xc000442618, 0x72, 0x0, 0x0, 0xbadadd)
          /usr/local/go/src/internal/poll/fd_poll_runtime.go:87 +0x9b
  internal/poll.(*pollDesc).waitRead(...)
          /usr/local/go/src/internal/poll/fd_poll_runtime.go:92
  internal/poll.(*FD).Accept(0xc000442600, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
          /usr/local/go/src/internal/poll/fd_unix.go:384 +0x1ba
  net.(*netFD).accept(0xc000442600, 0xb51080, 0x50, 0xc0000a77c0)
          /usr/local/go/src/net/fd_unix.go:238 +0x42
  net.(*TCPListener).accept(0xc00049e110, 0xc000046a80, 0x7fccf3be26d0, 0xc000000180)
          /usr/local/go/src/net/tcpsock_posix.go:139 +0x32
  net.(*TCPListener).AcceptTCP(0xc00049e110, 0x40d9b8, 0x30, 0xb51080)
          /usr/local/go/src/net/tcpsock.go:247 +0x48
  net/http.tcpKeepAliveListener.Accept(0xc00049e110, 0xb51080, 0xc0002d0e70, 0xadef20, 0x2294c70)
          /usr/local/go/src/net/http/server.go:3264 +0x2f
  net/http.(*Server).Serve(0xc0002d7d40, 0xcb8640, 0xc00049e110, 0x0, 0x0)
          /usr/local/go/src/net/http/server.go:2859 +0x22d
  net/http.(*Server).ListenAndServe(0xc0002d7d40, 0xc0002d7d40, 0xc000355ea8)
          /usr/local/go/src/net/http/server.go:2797 +0xe4
  net/http.ListenAndServe(...)
          /usr/local/go/src/net/http/server.go:3037
  github.com/gin-gonic/gin.(*Engine).Run(0xc000394000, 0xc000355f48, 0x1, 0x1, 0x0, 0x0)
          /root/go/src/github.com/duyanghao/gin-apiserver/vendor/github.com/gin-gonic/gin/gin.go:294 +0x140
  main.main()
          /root/go/src/github.com/duyanghao/gin-apiserver/cmd/main.go:22 +0x2c4
  
  goroutine 19 [syscall]:
  os/signal.signal_recv(0xcb28a0)
          /usr/local/go/src/runtime/sigqueue.go:139 +0x9c
  os/signal.loop()
          /usr/local/go/src/os/signal/signal_unix.go:23 +0x22
  created by os/signal.init.0
          /usr/local/go/src/os/signal/signal_unix.go:29 +0x41
  
  === END goroutine stack dump ===
  ```

## Framework

`gin-apiserver`框架的核心就是pkg包，下面主要针对该包结构进行描述：

```bash
pkg/
├── config
│   ├── config.go
│   ├── key.go
│   ├── model.go
│   └── opt_defs.go
├── controller
│   ├── ping.go
│   ├── todo.go
│   └── version.go
├── log
│   └── log.go
├── middleware
│   ├── basic_auth_middleware.go
├── models
│   └── common.go
├── route
│   └── routes.go
├── service
│   └── todo.go
├── store
└── util
```

* config：主要用于配置文件，实现：文件+环境变量+命令行参数读取
* controller: 对应MVC中controller，调用service中的接口进行实际处理，自己只做数据校验与拼接
* service: 负责主要的逻辑实现
* log: 日志模块，实现：模块名(文件名)+函数名+行数+日志级别
* middleware: 中间件，负责通用的处理，例如：鉴权
* models: 对应MVC中的model
* route: gin路由
* store: 存储模块，可以添加MySQL、Redis等
* util: 通用的库函数

## Usage

* step1 - 替换项目名称

  实际使用中，通常需要将`gin-apiserver`替换成业务需要的后台server名称，可以执行如下命令：

  ```bash
  $ grep -rl gin-apiserver . | xargs sed -i 's/gin-apiserver/youapiserver/g' 
  ```
  
* step2 - 开发业务controller和service

  框架中已经集成了一个示例(todo)：
  
  ```go
  // controller(pkg/controller/todo.go)
  type ToDoController interface {
  	GetToDo(c *gin.Context)
  }
  
  // service(pkg/service/todo.go)
  type ToDoService interface {
  	Get()
  }
  ```
  
  我们需要按照自身业务需求开发todo(替换成任意功能)的controller和service逻辑。另外你也可以参考todo添加其它功能对应的controller和service
   
* step3 - 启动服务  

  可以直接启动运行服务，如下：

  ```bash
  $ bash hack/start.sh
  ```
  
  也可以在Kubernetes集群中启动服务，如下：
  
  ```bash
  # generated image
  $ make dockerfiles.build
  # retag and push to your docker registry
  $ docker tag duyanghao/gin-apiserver:v0.1.0 xxx/duyanghao/gin-apiserver:v0.1.0
  $ docker push xxx/duyanghao/gin-apiserver:v0.1.0
  # Update the deployment to use the built image name
  $ sed -i 's|REPLACE_IMAGE|xxx/duyanghao/gin-apiserver:v0.1.0|g' hack/deploy/deployment.yaml
  # create service 
  $ kubectl apply -f hack/deploy/service.yaml
  # create deployment
  $ kubectl apply -f hack/deploy/deployment.yaml
  ```

## 3rd party Implements

* [coredns-dynapi-adapter - coredns dynamic middleware apiserver adapter](https://github.com/duyanghao/coredns-dynapi-adapter)

## Refs

* [dump-goroutine-stack-traces](https://colobu.com/2016/12/21/how-to-dump-goroutine-stack-traces/)
