package netxdcustomerdalservices

import (
	"context"
	netxdcustomerdalinterfaces "task-grpc/netxd_customer_dal/netxd-customer_dal_interfaces"
	netxdcustomerdalmodels "task-grpc/netxd_customer_dal/netxd_customer_dal_models"
	"time"

	//"errors"
	//"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) netxdcustomerdalinterfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (p *CustomerService) CreateCustomer(user *netxdcustomerdalmodels.Customer) (*netxdcustomerdalmodels.CustomerResponse, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.IsActive = false

	_, err := p.CustomerCollection.InsertOne(p.ctx, &user)
	if err != nil {
		return nil, err
	}
	response := &netxdcustomerdalmodels.CustomerResponse{
		CustomerId: 100,
	}

	return response, nil
}
