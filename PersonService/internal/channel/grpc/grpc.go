package grpc

import (
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type grpcChan struct {
	service service.Service
}

func New() *grpcChan {
	return &grpcChan{
		service: service.New(),
	}
}

func (g *grpcChan) Start() {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("Cannot open the grpc server")
	}

	server.Serve(listener)
}

func (r *grpcChan) Create(ctx context.Context, request *pb.Person) (*pb.Res, error) {

	person := canonical.Person{
		Id:         "",
		Name:       request.Name,
		SecondName: request.SecondName,
		Age:        request.Age,
		Document:   request.Document,
	}

	id, err := r.service.Create(person)
	if err != nil {
		return nil, err
	}

	return &pb.Res{
		Id: id,
	}, nil
}
