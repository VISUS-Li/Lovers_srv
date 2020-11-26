// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: homecard.proto

package lovers_srv_home

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Home service

type HomeService interface {
	GetMainCard(ctx context.Context, in *GetMainCardReq, opts ...client.CallOption) (*GetMainCardResp, error)
	GetCardByCount(ctx context.Context, in *GetCardByCountReq, opts ...client.CallOption) (*GetCardByCountResp, error)
	GetCardByIndex(ctx context.Context, in *GetCardByIndexReq, opts ...client.CallOption) (*GetCardByIndexResp, error)
	PostCardInfo(ctx context.Context, in *PostCardInfoReq, opts ...client.CallOption) (*PostCardInfoResp, error)
}

type homeService struct {
	c    client.Client
	name string
}

func NewHomeService(name string, c client.Client) HomeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "lovers.srv.home"
	}
	return &homeService{
		c:    c,
		name: name,
	}
}

func (c *homeService) GetMainCard(ctx context.Context, in *GetMainCardReq, opts ...client.CallOption) (*GetMainCardResp, error) {
	req := c.c.NewRequest(c.name, "Home.GetMainCard", in)
	out := new(GetMainCardResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeService) GetCardByCount(ctx context.Context, in *GetCardByCountReq, opts ...client.CallOption) (*GetCardByCountResp, error) {
	req := c.c.NewRequest(c.name, "Home.GetCardByCount", in)
	out := new(GetCardByCountResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeService) GetCardByIndex(ctx context.Context, in *GetCardByIndexReq, opts ...client.CallOption) (*GetCardByIndexResp, error) {
	req := c.c.NewRequest(c.name, "Home.GetCardByIndex", in)
	out := new(GetCardByIndexResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeService) PostCardInfo(ctx context.Context, in *PostCardInfoReq, opts ...client.CallOption) (*PostCardInfoResp, error) {
	req := c.c.NewRequest(c.name, "Home.PostCardInfo", in)
	out := new(PostCardInfoResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Home service

type HomeHandler interface {
	GetMainCard(context.Context, *GetMainCardReq, *GetMainCardResp) error
	GetCardByCount(context.Context, *GetCardByCountReq, *GetCardByCountResp) error
	GetCardByIndex(context.Context, *GetCardByIndexReq, *GetCardByIndexResp) error
	PostCardInfo(context.Context, *PostCardInfoReq, *PostCardInfoResp) error
}

func RegisterHomeHandler(s server.Server, hdlr HomeHandler, opts ...server.HandlerOption) error {
	type home interface {
		GetMainCard(ctx context.Context, in *GetMainCardReq, out *GetMainCardResp) error
		GetCardByCount(ctx context.Context, in *GetCardByCountReq, out *GetCardByCountResp) error
		GetCardByIndex(ctx context.Context, in *GetCardByIndexReq, out *GetCardByIndexResp) error
		PostCardInfo(ctx context.Context, in *PostCardInfoReq, out *PostCardInfoResp) error
	}
	type Home struct {
		home
	}
	h := &homeHandler{hdlr}
	return s.Handle(s.NewHandler(&Home{h}, opts...))
}

type homeHandler struct {
	HomeHandler
}

func (h *homeHandler) GetMainCard(ctx context.Context, in *GetMainCardReq, out *GetMainCardResp) error {
	return h.HomeHandler.GetMainCard(ctx, in, out)
}

func (h *homeHandler) GetCardByCount(ctx context.Context, in *GetCardByCountReq, out *GetCardByCountResp) error {
	return h.HomeHandler.GetCardByCount(ctx, in, out)
}

func (h *homeHandler) GetCardByIndex(ctx context.Context, in *GetCardByIndexReq, out *GetCardByIndexResp) error {
	return h.HomeHandler.GetCardByIndex(ctx, in, out)
}

func (h *homeHandler) PostCardInfo(ctx context.Context, in *PostCardInfoReq, out *PostCardInfoResp) error {
	return h.HomeHandler.PostCardInfo(ctx, in, out)
}