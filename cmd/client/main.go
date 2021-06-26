package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	proto "github.com/vinhlh/grpc-playground/order-service/v1"
)

func main() {
	conn, err := grpc.Dial("localhost:1106", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}

	defer conn.Close()

	c := proto.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := c.GetOrderById(ctx, &proto.Request{
		OrderId: "m3vl",
	})
	if err != nil {
		log.Fatalf("can't fetch order id: %v", err)
	}

	fmt.Println("found", r)
}
