package main

import (
	hello "Lovers_srv/proto/greeter"
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"strings"
)

type Say struct{
	Client hello.GreeterService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error{
	name,ok := req.Get["name"]
	if !ok || len(name.Values) == 0{
		return errors.BadRequest("go.micro.api.greeter","名字不能为空")
	}

	response, err :=s.Client.Hello(ctx, &hello.Request{
		Name:strings.Join(name.Values," "),
	})
	if err != nil{
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message":response.Greeting,
	})

	rsp.Body = string(b)
	return nil
}

func main(){
	reg := consul.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.api.greeter"),
	)
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client:hello.NewGreeterService("go.micro.srv.greeter",service.Client())}))
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}