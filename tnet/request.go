package tnet

import "github.com/ibuprofen/Tin/tinface"

type Request struct {
	conn tinface.IConnection
	msg  tinface.IMessage
}

func (r *Request) GetConnection() tinface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
