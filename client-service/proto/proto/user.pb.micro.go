// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error)
	UpdateUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error)
	SelectUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) CreateUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error) {
	req := c.c.NewRequest(c.name, "UserService.CreateUser", in)
	out := new(ResponseResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateUser", in)
	out := new(ResponseResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SelectUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error) {
	req := c.c.NewRequest(c.name, "UserService.SelectUser", in)
	out := new(ResponseResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*ResponseResult, error) {
	req := c.c.NewRequest(c.name, "UserService.DeleteUser", in)
	out := new(ResponseResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	CreateUser(context.Context, *UserRequest, *ResponseResult) error
	UpdateUser(context.Context, *UserRequest, *ResponseResult) error
	SelectUser(context.Context, *UserRequest, *ResponseResult) error
	DeleteUser(context.Context, *UserRequest, *ResponseResult) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		CreateUser(ctx context.Context, in *UserRequest, out *ResponseResult) error
		UpdateUser(ctx context.Context, in *UserRequest, out *ResponseResult) error
		SelectUser(ctx context.Context, in *UserRequest, out *ResponseResult) error
		DeleteUser(ctx context.Context, in *UserRequest, out *ResponseResult) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) CreateUser(ctx context.Context, in *UserRequest, out *ResponseResult) error {
	return h.UserServiceHandler.CreateUser(ctx, in, out)
}

func (h *userServiceHandler) UpdateUser(ctx context.Context, in *UserRequest, out *ResponseResult) error {
	return h.UserServiceHandler.UpdateUser(ctx, in, out)
}

func (h *userServiceHandler) SelectUser(ctx context.Context, in *UserRequest, out *ResponseResult) error {
	return h.UserServiceHandler.SelectUser(ctx, in, out)
}

func (h *userServiceHandler) DeleteUser(ctx context.Context, in *UserRequest, out *ResponseResult) error {
	return h.UserServiceHandler.DeleteUser(ctx, in, out)
}
