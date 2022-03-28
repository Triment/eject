### eject

一个快速web框架

+ [x] 中间件系统（单向工作流，非洋葱模型）
+ [x] 集成路由中间件

##### 内置路由特性


参数路由

|注册路由|实际路由|匹配参数|
|-----|-----|-----|
|/hello/:user/:post|/hello/a/b|`{ user: a, post: b }`|
|/public/*path|/public/file/a.txt|`{ path: file/a.txt }`

参数路由的优先级`:`优先于`*`，如下面两个路由注册后会将path将匹配为123

注册路由1: `/public/:path/hello`

注册路由2: `/public/*path`

实际路由: `/public/123/456`

解析后的map为
```json
{
    "path": "123"
}
```
如果没有注册路由1，将会解析为：
```json
{
    "path": "123/456"
}
```
用法

```bash
go get -u https://github.com/triment/eject
```

```golang
package main

import (
	"github.com/Triment/eject"
)

type ResMessage struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func auth(context *eject.Context) bool {//鉴权中间件通过返回true来允许该路由正常运行，可在鉴权路由中进行响应
    if context.Req.Header.Get("user")!= "admin" {
    	return false
    }
    return true
}

func main(){
    router := eject.CreateRouter()//创建路由
    router.GET("/hello", func(context *eject.Context) {
        context.Res.Write([]byte("hello word")
     })//注册路由
    router.GET("/", func(context *eject.Context) {
        context.JSON(&ResMessage{Status: 200, Body: "请求成功"})
    })
    router.GET("/auth", func(context *eject.Context){
        context.JSON(&ResMessage{Status: 200, Body: "请求成功"})
    }).Before(auth)//鉴权中间件
    app := eject.CreateApp()//创建应用
	app.Inject(router.Accept())//注入路由中间件
	app.Listen(":4567")//监听端口 
}
```

TodoList
+ [ ] ~~目前实现的中间件是单向流，如同一个处理的大函数，只不过响应和路由分离为内置的路由中间件了，下一步实现面向切面，代码注入的形式~~
+ [x] ~~目前实现的请求路由仅支持GET， POST， 将来可能会加入更多请求方法，标准化框架~~
+ [ ] 分布式应用支持，利用TCP/UDP交换节点信息，自动负载均衡（待评估，需要用概率对实际应用分析后实现相应算法）

实际应用

* [fileserver](https://github.com/triment/fileserver.git) 一个文件下载系统，仅支持下载，上传请使用其他方式。
