package service

import (
	"context"
	"myapp/api/proto/myapp/api"
	"myapp/internal/biz"
)

type Hello struct {
	biz.HelloRequest
}

func NewHello(request biz.HelloRequest) *Hello {
	return &Hello{request}
}

func (h *Hello) SayHello(ctx context.Context, request *api.HelloRequest) (*api.HelloResponse, error) {
	bizRes := h.BizHello(ctx, biz.NewHelloReq(request.Name))
	return &api.HelloResponse{Message: bizRes.GetMessage()}, nil
}
