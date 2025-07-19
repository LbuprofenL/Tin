package tnet

import (
	"fmt"
	"strconv"

	"github.com/ibuprofen/Tin/tinface"
)

type MsgHandler struct {
	Apis map[uint32]tinface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]tinface.IRouter),
	}
}

func (mh *MsgHandler) DoMsgHandler(req tinface.IRequest) {
	rt, ok := mh.Apis[req.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", req.GetMsgID(), " is not FOUND!")
		return
	}
	rt.PreHandle(req)
	rt.Handle(req)
	rt.PostHandle(req)
}

func (mh *MsgHandler) AddRouter(msgID uint32, router tinface.IRouter) {
	// 1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Apis[msgID]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgID)))
	}
	// 2 添加msg与api的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api msgId = ", msgID)
}
