package gface

/*
消息管理器的抽象层
这里面有两个方法，AddRouter()就是添加一个msgId和一个路由关系到Apis中，
那么DoMsgHandler()则是调用Router中具体Handle()等方法的接口。
*/

type IMsgHandle interface {
	DoMsgHandler(request IRequest)          //马上以非阻塞方式处理消息
	AddRouter(msgId uint32, router IRouter) //为消息添加具体的处理逻辑
}
