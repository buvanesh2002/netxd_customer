package netxdcustomerdalinterfaces

import netxdcustomerdalmodels "task-grpc/netxd_customer_dal/netxd_customer_dal_models"



type ICustomer interface {
	CreateCustomer(user *netxdcustomerdalmodels.Customer) (*netxdcustomerdalmodels.CustomerResponse, error)
}
