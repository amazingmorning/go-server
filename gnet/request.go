package gnet

import "go-server/gface"

type Request struct {
	conn gface.IConnection //已经和客户端建立的链接
	msg  gface.IMessage    //客户端请求的数据
}

//获取请求链接的信息
func (r *Request) GetConnection() gface.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//获取请求的消息ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
