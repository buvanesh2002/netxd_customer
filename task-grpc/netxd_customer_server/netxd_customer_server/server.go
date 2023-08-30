package main

import (
	"context"
	"fmt"
	"net"
	netxdcustomerdalservices "task-grpc/netxd_customer_dal/netxd_customer_dal_services"
	netxdcustomerserverconfig "task-grpc/netxd_customer_server/netxd_customer_server_config"
	netxdcustomerserverconstants "task-grpc/netxd_customer_server/netxd_customer_server_constants"
	netxdcustomerservercontroller "task-grpc/netxd_customer_server/netxd_customer_server_controller"

	pro "task-grpc/netxd_customer/netxd_customer"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	profileCollection := netxdcustomerserverconfig.GetCollection(client, "bankdb", "profile")
	netxdcustomerservercontroller.CustomerService = netxdcustomerdalservices.InitCustomerService(profileCollection, context.Background())
}

func main() {
	mongoclient, err := netxdcustomerserverconfig.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", netxdcustomerserverconstants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &netxdcustomerservercontroller.RPCServer{})

	fmt.Println("Server listening on", netxdcustomerserverconstants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
