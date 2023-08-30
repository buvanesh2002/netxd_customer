package netxdcustomerservercontroller

import (
	"context"
	netxdcustomerdalinterfaces "task-grpc/netxd_customer_dal/netxd-customer_dal_interfaces"
	netxdcustomerdalmodels "task-grpc/netxd_customer_dal/netxd_customer_dal_models"
	pro "task-grpc/netxd_customer/netxd_customer"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService netxdcustomerdalinterfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {
	dbCustomer:= &netxdcustomerdalmodels.Customer{
		CustomerId: req.CustomerId,
		FirstName: req.FirstName,
		LastName: req.LastName,
		BankId: req.BankId,
		Balance: float64(req.Balance),
	
	}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseProfile := &pro.CustomerResponse{
			CustomerId: (result.CustomerId),
		}
		return responseProfile, nil
	}
}