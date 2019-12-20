## config 

1、首先从根目录下的config.json文件加载配置

2、使用系统参数覆盖配置，这样使用k8s部署时可以灵活设置参数

## controller

### 接收前端请求

#### 获取get参数

打开 index.go 文件

获取连接 http://localhost:8081/api/v1/test/index?par1=测试 中par1的值

```
// Index 首页
func Index(writer http.ResponseWriter, request *http.Request) {
  request.ParseForm()
  par1:= getParam("par1",request)
}
func getParam(parameterName string, request *http.Request) string {
	if len(request.Form[parameterName]) > 0 {
		return request.Form[parameterName][0]
	}
	return ""
}
```
#### 获取POST参数

打开 index.go 文件

```
// GetBody2Struct 获取POST参数，并转化成指定的struct对象
func GetBody2Struct(request *http.Request, pojo interface{}) error {
	s, _ := ioutil.ReadAll(request.Body)
	if len(s) == 0 {
		return nil
	}
	return json.Unmarshal(s, pojo)
}
```

## model

### db 连接

打开database.go文件

#### 新建连接

```
db, err = gorm.Open(conf.DbType, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
```

#### 自动新建表格

```
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").
		AutoMigrate(&Test{})
```

#### 添加索引

```
db.Model(&Test{}).AddIndex("idx_id", "id")
```

## router 路由

打开 router.go 文件

### 路由设置

```
Mux.HandleFunc("/api/v1/test/index", interceptor(controller.Index))
```

### 跨域设置

```
func crossOrigin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", conf.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", conf.AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", conf.AccessControlAllowHeaders)
		h(w, r)
	}
}
```


## signal 信号设置

主要用于接收应用异常关闭信号，如Ctrl+C等，接收到信号之后关闭所有启动的资源和goroutine防止内存泄露

打开 signal.go 文件

```
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

// shutdownRequestChannel 用于关闭初始化
var shutdownRequestChannel = make(chan struct{})

// interruptSignals 定义默认的关闭触发信号
var interruptSignals = []os.Signal{os.Interrupt}

// interruptListener 监听关闭信号(Ctrl+C)
func interruptListener(s *http.Server) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		interruptChannel := make(chan os.Signal, 1)
		signal.Notify(interruptChannel, interruptSignals...)
		select {
		case sig := <-interruptChannel:
			log.Printf("收到关闭信号 (%s). 关闭...\n", sig)
			s.Close()
		case <-shutdownRequestChannel:
			log.Println("关闭请求.关闭...")
			s.Close()
		}
		close(c)
		// 重复关闭信号处理
		for {
			select {
			case sig := <-interruptChannel:
				log.Printf("Received signal (%s).  Already "+
					"shutting down..\n", sig)
			case <-shutdownRequestChannel:
				log.Println("Shutdown requested.  Already " +
					"shutting down...")
			}
		}
	}()
	return c
}
func interruptRequested(interrupted <-chan struct{}) bool {
	select {
	case <-interrupted:
		return true
	default:
	}
	return false
}

```

## main

打开 main.go文件




