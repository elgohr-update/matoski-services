// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/account.proto

package account

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

// Client API for Account service

type AccountService interface {
	Login(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*Response, error)
	UpdateProfilePicture(ctx context.Context, in *Picture, opts ...client.CallOption) (*Response, error)
	Health(ctx context.Context, in *Healthcheck, opts ...client.CallOption) (*Healthcheck, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "account"
	}
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) Login(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Account.Login", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Account.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Account.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Account.UpdatePassword", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) UpdateProfilePicture(ctx context.Context, in *Picture, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Account.UpdateProfilePicture", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Health(ctx context.Context, in *Healthcheck, opts ...client.CallOption) (*Healthcheck, error) {
	req := c.c.NewRequest(c.name, "Account.Health", in)
	out := new(Healthcheck)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Login(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	Update(context.Context, *User, *Response) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *Response) error
	UpdateProfilePicture(context.Context, *Picture, *Response) error
	Health(context.Context, *Healthcheck, *Healthcheck) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) error {
	type account interface {
		Login(ctx context.Context, in *User, out *Response) error
		Get(ctx context.Context, in *User, out *Response) error
		Update(ctx context.Context, in *User, out *Response) error
		UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *Response) error
		UpdateProfilePicture(ctx context.Context, in *Picture, out *Response) error
		Health(ctx context.Context, in *Healthcheck, out *Healthcheck) error
	}
	type Account struct {
		account
	}
	h := &accountHandler{hdlr}
	return s.Handle(s.NewHandler(&Account{h}, opts...))
}

type accountHandler struct {
	AccountHandler
}

func (h *accountHandler) Login(ctx context.Context, in *User, out *Response) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *accountHandler) Get(ctx context.Context, in *User, out *Response) error {
	return h.AccountHandler.Get(ctx, in, out)
}

func (h *accountHandler) Update(ctx context.Context, in *User, out *Response) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *accountHandler) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *Response) error {
	return h.AccountHandler.UpdatePassword(ctx, in, out)
}

func (h *accountHandler) UpdateProfilePicture(ctx context.Context, in *Picture, out *Response) error {
	return h.AccountHandler.UpdateProfilePicture(ctx, in, out)
}

func (h *accountHandler) Health(ctx context.Context, in *Healthcheck, out *Healthcheck) error {
	return h.AccountHandler.Health(ctx, in, out)
}
