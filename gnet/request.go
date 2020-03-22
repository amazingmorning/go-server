package gnet

import "go-server/gface"

type Request struct {
	conn gface.IConnection //已经和客户端建立的链接
	data []byte            //客户端请求的数据
}

//获取请求链接的信息
func (r *Request) GetConnection() gface.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.data
}
