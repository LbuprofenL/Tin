package tnet

import "github.com/ibuprofen/Tin/tin/tinface"

type Request struct {
	conn tinface.IConnection
	data []byte
}

func (r *Request) GetConnection() tinface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
