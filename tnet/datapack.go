package tnet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/ibuprofen/Tin/tinface"
	"github.com/ibuprofen/Tin/utils"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// ID uint32(4byte) + Datalen uint32(4byte)
	return 8
}

func (dp *DataPack) Pack(msg tinface.IMessage) ([]byte, error) {
	dataBuf := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuf.Bytes(), nil
}

func (dp *DataPack) Unpack(data []byte) (tinface.IMessage, error) {
	dataBuf := bytes.NewReader(data)

	msg := &Message{}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("Too large msg data recieved")
	}
	// 这里只拆了head数据，data数据需要再读取一次
	return msg, nil
}
