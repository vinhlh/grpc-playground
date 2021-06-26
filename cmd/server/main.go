package main

import (
	"context"
	"log"
	"net"

	proto "github.com/vinhlh/grpc-playground/order-service/v1"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedOrderServiceServer
}

func (s *server) GetOrderById(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	return &proto.Response{
		Id: req.OrderId,
		Status: proto.Status_ORDER_COMPLETED,
		Customer: &proto.Customer{
			FirstName: "Vinh",
			LastName:  "Le",
			Email:     "hung.vjnh@gmail.com",
		},
		Products: []*proto.Product{
			{
				Id:    "777",
				Name:  "Nike Running Tee",
				Price: "10.0",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:1106")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
