// Code generated by goctl. DO NOT EDIT!
// Source: hi.proto

package client

import (
	"context"

	"github.com/Holyson/test-go-zero-cors/tools/goctl/example/rpc/hi/pb/hi"

	"github.com/Holyson/test-go-zero-cors/zrpc"
	"google.golang.org/grpc"
)

type (
	EventReq  = hi.EventReq
	EventResp = hi.EventResp
	HelloReq  = hi.HelloReq
	HelloResp = hi.HelloResp
	HiReq     = hi.HiReq
	HiResp    = hi.HiResp

	Event interface {
		AskQuestion(ctx context.Context, in *EventReq, opts ...grpc.CallOption) (*EventResp, error)
	}

	defaultEvent struct {
		cli zrpc.Client
	}
)

func NewEvent(cli zrpc.Client) Event {
	return &defaultEvent{
		cli: cli,
	}
}

func (m *defaultEvent) AskQuestion(ctx context.Context, in *EventReq, opts ...grpc.CallOption) (*EventResp, error) {
	client := hi.NewEventClient(m.cli.Conn())
	return client.AskQuestion(ctx, in, opts...)
}
