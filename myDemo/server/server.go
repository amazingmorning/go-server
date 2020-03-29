package main

import (
	"fmt"
	"go-server/gface"
	"go-server/gnet"
)

//ping test 自定义路由
type PingRouter struct {
	gnet.BaseRouter
}

//Ping Handle
func (this *PingRouter) Handle(request gface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

//HelloZinxRouter Handle
type HelloZinxRouter struct {
	gnet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request gface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//创建一个server句柄
	s := gnet.NewServer()

	//配置路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//开启服务
	s.Serve()
}
