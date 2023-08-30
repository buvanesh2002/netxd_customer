package main

import (
	"context"
	"fmt"
	"log"

	pb "task-grpc/netxd_customer/netxd_customer"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.CustomerDetails{
		CustomerId: 2545,
		FirstName: "sanjay",
		LastName: "J",
		BankId: "987",
		Balance: 6500.00,
  	})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %d\n", response.CustomerId)
}
