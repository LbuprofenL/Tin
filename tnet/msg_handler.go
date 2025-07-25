package tnet

import (
	"fmt"
	"strconv"

	"github.com/ibuprofen/Tin/tinface"
	"github.com/ibuprofen/Tin/utils"
)

type MsgHandler struct {
	Apis           map[uint32]tinface.IRouter
	WorkerPoolSize uint32
	TaskQueue      []chan tinface.IRequest
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]tinface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan tinface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan tinface.IRequest, 4096)
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

func (mh *MsgHandler) StartOneWorker(workerID int, taskQueue chan tinface.IRequest) {
	fmt.Println("[Worker] Worker ID = ", workerID, " is started.")
	for {
		select {
		case req := <-taskQueue:
			mh.DoMsgHandler(req)
		}
	}
}

func (mh *MsgHandler) SendMsgToTaskQueue(req tinface.IRequest) {
	workerID := req.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID=", req.GetConnection().GetConnID(), " request msgID=", req.GetMsgID(), " to workerID=", workerID)
	mh.TaskQueue[workerID] <- req
}
