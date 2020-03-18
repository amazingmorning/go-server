package gface

type gface interface{
	//启动服务
	Start()
	//运行服务
	Serve()
	//终止服务
	Stop()
}