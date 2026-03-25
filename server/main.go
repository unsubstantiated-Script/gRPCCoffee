package main

import (
	"context"
	pb "gRPCCoffee/gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(ctx context.Context, req *pb.MenuRequest) (*pb.Menu, error) {
	return &pb.Menu{
		Items: []*pb.Item{
			{Id: "1", Name: "Black Coffee"},
			{Id: "2", Name: "Americano"},
			{Id: "3", Name: "Vanilla Soy Chai Latte"},
		},
	}, nil
}

func (s *server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{Id: "ABC123"}, nil
}

func (s *server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "In Progress",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	log.Println("Server started")
	log.Fatal(grpcServer.Serve(lis))
}
