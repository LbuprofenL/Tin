package tnet

import "github.com/ibuprofen/Tin/tin/tinface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request tinface.IRequest)  {}
func (br *BaseRouter) Handle(request tinface.IRequest)     {}
func (br *BaseRouter) PostHandle(request tinface.IRequest) {}
