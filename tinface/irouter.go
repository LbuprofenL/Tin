package tinface

type IRouter interface {
	PreHandle(request IRequest)
	Handle(request IRequest)
	PostHandle(request IRequest)
}
