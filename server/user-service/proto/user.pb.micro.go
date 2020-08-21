// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package lovers_srv_user

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

// Client API for User service

type UserService interface {
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...client.CallOption) (*LogoutResp, error)
	RegisterUser(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*RegisterResp, error)
	BindLover(ctx context.Context, in *BindLoverReq, opts ...client.CallOption) (*BindLoverResp, error)
	UnBindLover(ctx context.Context, in *UnBindLoverReq, opts ...client.CallOption) (*UnBindLoverResp, error)
	GetLoverInfo(ctx context.Context, in *GetLoverInfoReq, opts ...client.CallOption) (*GetLoverInfoResp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "lovers.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(LoginResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Logout(ctx context.Context, in *LogoutReq, opts ...client.CallOption) (*LogoutResp, error) {
	req := c.c.NewRequest(c.name, "User.Logout", in)
	out := new(LogoutResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RegisterUser(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*RegisterResp, error) {
	req := c.c.NewRequest(c.name, "User.RegisterUser", in)
	out := new(RegisterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) BindLover(ctx context.Context, in *BindLoverReq, opts ...client.CallOption) (*BindLoverResp, error) {
	req := c.c.NewRequest(c.name, "User.BindLover", in)
	out := new(BindLoverResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UnBindLover(ctx context.Context, in *UnBindLoverReq, opts ...client.CallOption) (*UnBindLoverResp, error) {
	req := c.c.NewRequest(c.name, "User.UnBindLover", in)
	out := new(UnBindLoverResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetLoverInfo(ctx context.Context, in *GetLoverInfoReq, opts ...client.CallOption) (*GetLoverInfoResp, error) {
	req := c.c.NewRequest(c.name, "User.GetLoverInfo", in)
	out := new(GetLoverInfoResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Login(context.Context, *LoginReq, *LoginResp) error
	Logout(context.Context, *LogoutReq, *LogoutResp) error
	RegisterUser(context.Context, *RegisterReq, *RegisterResp) error
	BindLover(context.Context, *BindLoverReq, *BindLoverResp) error
	UnBindLover(context.Context, *UnBindLoverReq, *UnBindLoverResp) error
	GetLoverInfo(context.Context, *GetLoverInfoReq, *GetLoverInfoResp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Login(ctx context.Context, in *LoginReq, out *LoginResp) error
		Logout(ctx context.Context, in *LogoutReq, out *LogoutResp) error
		RegisterUser(ctx context.Context, in *RegisterReq, out *RegisterResp) error
		BindLover(ctx context.Context, in *BindLoverReq, out *BindLoverResp) error
		UnBindLover(ctx context.Context, in *UnBindLoverReq, out *UnBindLoverResp) error
		GetLoverInfo(ctx context.Context, in *GetLoverInfoReq, out *GetLoverInfoResp) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Login(ctx context.Context, in *LoginReq, out *LoginResp) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) Logout(ctx context.Context, in *LogoutReq, out *LogoutResp) error {
	return h.UserHandler.Logout(ctx, in, out)
}

func (h *userHandler) RegisterUser(ctx context.Context, in *RegisterReq, out *RegisterResp) error {
	return h.UserHandler.RegisterUser(ctx, in, out)
}

func (h *userHandler) BindLover(ctx context.Context, in *BindLoverReq, out *BindLoverResp) error {
	return h.UserHandler.BindLover(ctx, in, out)
}

func (h *userHandler) UnBindLover(ctx context.Context, in *UnBindLoverReq, out *UnBindLoverResp) error {
	return h.UserHandler.UnBindLover(ctx, in, out)
}

func (h *userHandler) GetLoverInfo(ctx context.Context, in *GetLoverInfoReq, out *GetLoverInfoResp) error {
	return h.UserHandler.GetLoverInfo(ctx, in, out)
}