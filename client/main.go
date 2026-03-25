package main

import (
	"context"
	pb "gRPCCoffee/gen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	menu, err := client.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Menu: %v\n", menu.Items)

	reciept, err := client.PlaceOrder(ctx, &pb.Order{Items: menu.Items})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reciept: %v\n", reciept)

	status, err := client.GetOrderStatus(ctx, &pb.Receipt{Id: reciept.Id})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %v\n", status)
}
