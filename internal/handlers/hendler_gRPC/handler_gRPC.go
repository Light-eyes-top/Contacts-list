package hendler_gRPC

import (
	someRestApi "SomeRestApi"
	"SomeRestApi/internal/handlers/hendler_gRPC/pb"
	"SomeRestApi/internal/service"
	"context"
)

type ContactHandler struct {
	pb.UnimplementedContactServiceServer
	service service.Services
}

func (c *ContactHandler) Create(_ context.Context, req *pb.Contact) (*pb.Id, error) {
	id, err := c.service.CreateContact(*someRestApi.NewContactFromPb(req))
	return &pb.Id{Id: id}, err
}
func (c *ContactHandler) Get(_ context.Context, req *pb.Id) (*pb.Contact, error) {
	contact, err := c.service.GetContact(req.Id)
	return someRestApi.NewContactToPb(&contact), err
}
func (c *ContactHandler) Update(_ context.Context, req *pb.UpdateRequest) (*pb.Contact, error) {
	contact, err := c.service.UpdateContact(req.Id, *someRestApi.NewContactFromPb(req.Contact))
	return someRestApi.NewContactToPb(&contact), err
}
func (c *ContactHandler) Delete(_ context.Context, req *pb.Id) (*pb.Id, error) {
	id, err := c.service.DeleteContact(req.Id)
	return &pb.Id{Id: id}, err
}

func NewContactHendler(service *service.Service) *ContactHandler {
	return &ContactHandler{service: service}
}
